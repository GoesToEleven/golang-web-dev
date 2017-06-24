package main

import (
	"net/http"
	"text/template"
)

type person struct {
	Name  string
	Prefs []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{"bond", []string{"martinis", "shaken not stirred", "fast cars", "intelligence"}}
	p2 := person{"moneypenny", []string{"bond"}}
	p3 := person{"Q", []string{"gadgets", "bond"}}
	xp := []person{p1, p2, p3}

	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
