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
	http.HandleFunc("/", idx)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./served"))))
	http.ListenAndServe(":80", nil)
}

func idx(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
