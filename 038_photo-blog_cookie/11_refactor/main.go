package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value) {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	m := Model(cookie.Value)
	tpl.ExecuteTemplate(res, "index.html", m)
}
