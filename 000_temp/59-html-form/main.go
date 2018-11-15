package main

import (
	"html/template"
	"net/http"
)

type GData struct {
	Title string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/process", process)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "Acme Inc",
	}
	tpl.ExecuteTemplate(w, "index.gohtml", gd)
}

func about(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "ABOUT",
	}
	tpl.ExecuteTemplate(w, "about.gohtml", gd)
}

func contact(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "CONTACT",
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", gd)
}

func signup(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "SIGNUP",
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", gd)
}

func process(w http.ResponseWriter, r *http.Request) {

	fn := r.FormValue("first")
	ln := r.FormValue("last")

	d := struct {
		GData
		First string
		Last  string
	}{
		GData: GData{
			Title: "PROCESS",
		},
		First: fn,
		Last:  ln,
	}

	tpl.ExecuteTemplate(w, "process.gohtml", d)
}
