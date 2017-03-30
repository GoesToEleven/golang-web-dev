package main

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		u := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: u.String(),
			Path: "/",
			HttpOnly: true,
		}
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}
