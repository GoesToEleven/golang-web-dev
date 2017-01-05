package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type person struct {
	First      string
	Last       string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribed") == "on"

	p1 := person{
		f,
		l,
		s,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", p1)

}
