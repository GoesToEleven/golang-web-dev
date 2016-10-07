package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, rq *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "index.gohtml", 42)
}
