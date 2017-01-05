package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are at foo")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are at bar")
}
