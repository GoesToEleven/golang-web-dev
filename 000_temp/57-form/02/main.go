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
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")
	surf := r.FormValue("surf")
	snow := r.FormValue("snow")
	skate := r.FormValue("skate")
	radio := r.FormValue("cow")

	d := struct {
		First, Last, Surf, Snow, Skate, Radio string
	}{
		First: fname,
		Last:  lname,
		Surf:  surf,
		Snow:  snow,
		Skate: skate,
		Radio: radio,
	}

	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
