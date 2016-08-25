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

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.ListenAndServe(":8080", nil)
}
