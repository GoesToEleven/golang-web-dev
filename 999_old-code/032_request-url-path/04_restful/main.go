package main

import (
	"io"
	"net/http"
	"strings"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}
	io.WriteString(res, `
	Dog Name: <strong>`+dogName+`</strong><br>
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {
	var dog DogHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)

	http.ListenAndServe(":9000", mux)
}
