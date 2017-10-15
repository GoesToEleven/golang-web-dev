package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml", "Snowykite")
}


func about(w http.ResponseWriter, r *http.Request){
	type customData struct {
		Title string
		Members []string
	}

	cd := customData {
		Title: "What I learned at SVCC:",
		Members: []string{"How to build a go server","How to serve files on go server", "How to use template", "Go is awesome!",},
	}

	tpl.ExecuteTemplate(w, "about.gohtml", cd)
}
