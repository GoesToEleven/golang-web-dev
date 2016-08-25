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
	http.ServeFile(res, req, "dog.jpeg")
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/toby", chien)
	http.ListenAndServe(":8080", nil)
}
