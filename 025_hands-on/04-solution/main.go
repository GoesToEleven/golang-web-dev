package main

import (
	"io"
	"net/http"
)

type snoop int

func (h snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {
	var dog snoop

	mux := http.NewServeMux()
	mux.Handle("/", dog)

	http.ListenAndServe(":8080", mux)
}
