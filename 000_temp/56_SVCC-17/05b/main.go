package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "ACME INC")
}

func about(w http.ResponseWriter, r *http.Request) {
	type customData struct {
		Title   string
		Members []string
	}

	cd := customData{
		Title:   "THE TEAM",
		Members: []string{"Moneypenny", "Bond", "Q", "M"},
	}

	tpl.ExecuteTemplate(w, "about.gohtml", cd)
}
