package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", dogPic)
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
	Dog Name: <strong>`+dogName+`</strong><br>
	<img src="/toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}