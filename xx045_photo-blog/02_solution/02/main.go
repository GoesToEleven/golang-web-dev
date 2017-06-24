package main

import (
	"github.com/GoesToEleven/golang-web-dev/045_photo-blog/02_solution/02/controllers"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template
var ctl *controllers.Controller

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	ctl = controllers.NewController(tpl)
}

func main() {
	r := httprouter.New()
	r.GET("/", ctl.Index)
	r.POST("/", ctl.IndexSubmission)
	r.Handler("GET", "/favicon.ico", http.NotFoundHandler())
	r.ServeFiles("/public/*filepath", http.Dir("./public"))
	http.ListenAndServe(":8080", r)
}
