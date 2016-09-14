package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func serveTemplate(res http.ResponseWriter, templateName string, params interface{}) {
	err := tpl.ExecuteTemplate(res, templateName, &params)
	HandleError(res, err)
}