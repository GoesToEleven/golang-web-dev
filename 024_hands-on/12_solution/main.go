package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
