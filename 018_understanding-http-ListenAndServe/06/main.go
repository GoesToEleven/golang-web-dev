package main

import (
	"io"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, "<h1>kitty kitty kitty<h1>")
	case "/dog":
		io.WriteString(res, "<h1>doggy doggy doggy<h1>")
	}
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
