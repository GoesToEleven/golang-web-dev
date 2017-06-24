package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type person struct {
	First string
	Age   int
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		"Todd",
		45,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", p1)
}
