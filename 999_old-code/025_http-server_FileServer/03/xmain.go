package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}

func chien(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<img src="/resources/dog.jpeg">
	<img src="/resources/dog1.jpeg">
	<img src="/resources/dog2.jpeg">
	`)
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/toby", chien)

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}
