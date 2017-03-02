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
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/css/main.css")
}

func bali(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/img/bali.jpg")
}