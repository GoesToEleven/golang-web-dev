package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	// we convert a value of type func(ResponseWriter, *Request) to HandlerFunc
	http.Handle("/", http.HandlerFunc(foo)) // conversion; handle takes a handler
	http.Handle("/dog/", http.HandlerFunc(bar)) // conversion; handle takes a handler
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, 4)
}
