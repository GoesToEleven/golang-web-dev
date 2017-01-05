package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/dog", bar)
	mux.HandleFunc("/cat", barred)
	http.ListenAndServe(":8080", mux)
}

func foo(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", 42)
	if err != nil {
		log.Println(err)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", 7)
	if err != nil {
		log.Println(err)
	}
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, catttty!")
}
