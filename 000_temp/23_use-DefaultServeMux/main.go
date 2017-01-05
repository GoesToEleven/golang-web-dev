package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var h hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", h)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}

// change code to use DefaultServeMux
// add a route for hamsters
