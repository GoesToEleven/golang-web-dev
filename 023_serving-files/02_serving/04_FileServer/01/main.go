package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/uptown/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	var dogName string

	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}

	io.WriteString(res, `
	<h1>Dog Name:`+dogName+`</h1>
	<img src="/toby.jpg">
	`)
}