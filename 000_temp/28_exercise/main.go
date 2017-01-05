package main

import (
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", foo)
	http.Handle("/cat", bar)

	http.ListenAndServe(":8080", nil)
}
