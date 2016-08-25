package main

import (
	"io"
	"net/http"
	"strings"
)

type Snoop int

func (h Snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path, "/")
	s := strings.Join(fs, " - ")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+s+`</h1><br>`)
}

func main() {
	var dog Snoop

	mux := http.NewServeMux()
	mux.Handle("/", dog)

	http.ListenAndServe(":8080", mux)
}
