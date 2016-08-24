package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.go"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/420", stoners)
	http.HandleFunc("/social", media)
	http.HandleFunc("/communities", expectancy)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Redirect(res, req, "/", 303)
	}
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func stoners(res http.ResponseWriter, req *http.Request) {
	type stats struct {
		data [][]string
	}
}

func media(res http.ResponseWriter, req *http.Request) {}

func expectancy(res http.ResponseWriter, req *http.Request) {}
