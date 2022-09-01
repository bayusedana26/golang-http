package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Writer string `json: "writer"`
	Title  string `json: "title`
}

type Articles []Article

var data = Articles{
	Article{Writer: "Bayu Sedana", Title: "Book of Fire"},
	Article{Writer: "Gale", Title: "Book of Air"},
}

func main() {
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/articles", GetArticles)
	http.HandleFunc("/post-article", withLogging(PostArticle))
	http.ListenAndServe(":3000", nil)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Youre in Home"))
}

func GetArticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(data)
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var newData Article

		err := json.NewDecoder(r.Body).Decode(&newData)

		if err != nil {
			fmt.Println("Error", err)
		}

		data = append(data, newData)

		json.NewEncoder(w).Encode(data)

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Can't read body", http.StatusInternalServerError)
		// }
		// w.Write([]byte(string(body)))

	} else {
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Log connection from", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
