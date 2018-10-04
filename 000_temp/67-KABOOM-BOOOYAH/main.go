package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.gohtml", nil)
}

func about (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func contact (w http.ResponseWriter, r * http.Request) {
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)
}
