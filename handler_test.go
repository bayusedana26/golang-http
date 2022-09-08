package golanghttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Word")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "This is home") })
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hi there") })
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "This images") })
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "This thumnbnails") })

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: mux,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.URL)
	}

	server := http.Server{
		Addr:    "localhost:8282",
		Handler: handler,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
