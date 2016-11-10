package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are at foo")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are at bar")
}
