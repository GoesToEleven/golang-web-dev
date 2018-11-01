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
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", bar )
	http.HandleFunc("/contact", con)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", 42)
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", []int{1,2,3,4,5,6,7,8,9})
}

func con(w http.ResponseWriter, r *http.Request) {

	type person struct {
		First string
		Age int
	}

	p1 := person{
		First: "James",
		Age: 32,
	}

	p2 := person {
		First: "Jenny",
		Age: 27,
	}

	xp := []person{p1, p2}


	tpl.ExecuteTemplate(w, "contact.gohtml", xp)
}
