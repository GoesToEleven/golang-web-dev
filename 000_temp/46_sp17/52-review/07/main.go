package main

import (
	"html/template"
	"net/http"
)

type person struct {
	First  string
	Last   string
	Saying string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First:  "James",
		Last:   "Bond",
		Saying: "Shaken, not stirred.",
	}

	p2 := person{
		First:  "Miss",
		Last:   "Moneypenny",
		Saying: "Hello, James. It is sooooooooo good to see you.",
	}

	xp := []person{p1, p2}
	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}
