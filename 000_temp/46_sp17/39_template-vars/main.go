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
	http.HandleFunc("/vars", vars)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func vars(w http.ResponseWriter, r *http.Request) {
	xs := []string{"James", "Moneypenny", "Q", "M"}
	tpl.ExecuteTemplate(w, "vars.gohtml", xs)
}
