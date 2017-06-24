package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar/", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from index")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from bar")
}
