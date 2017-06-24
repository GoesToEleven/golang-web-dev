package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return // comment out and experiment to see what happens
	http.Redirect(w, r, "/bar", http.StatusSeeOther)
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "bar.gohtml", nil)
}
