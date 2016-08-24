package main

import (
	"io"
	"net/http"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}

func main() {
	var dog DogHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)

	http.ListenAndServe(":8080", mux)
}

// https://godoc.org/net/http#ResponseWriter
// https://godoc.org/net/http#Header
// https://godoc.org/net/http#Header.Set

// https://en.wikipedia.org/wiki/List_of_HTTP_header_fields
