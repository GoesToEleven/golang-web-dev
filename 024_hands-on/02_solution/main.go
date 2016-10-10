package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/uptown/", dog)
	http.HandleFunc("/fido/", chien)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("fido.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}

func chien(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "dog.jpeg")
}
