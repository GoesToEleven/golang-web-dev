package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func serveTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := tpl.ExecuteTemplate(w, templateName, data)
	HandleError(w, err)
}
