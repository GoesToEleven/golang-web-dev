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
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/bar/", bar)
	mux.HandleFunc("/about", about)
	http.ListenAndServe(":8080", mux)
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "foo.gohtml", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "bar.gohtml", req.URL.Path)
}

func about(w http.ResponseWriter, req *http.Request) {

	d := struct {
		FName    string
		LName    string
		RawQuery string
	}{
		"James",
		"Bond",
		req.URL.RawQuery,
	}

	tpl.ExecuteTemplate(w, "about.gohtml", d)
}
