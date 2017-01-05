package main

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/GoesToEleven/golang-web-dev/045_photo-blog/02_solution/02/controllers"
	"google.golang.org/appengine"
)

var tpl *template.Template
var ctl *controllers.Controller

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	ctl = controllers.NewController(tpl)
}

func main() {
	// setup julienschmidt's router
	// julienschmidt's router implements handler interface
	r := httprouter.New()
	http.Handle("/", r)
	// user julienschmidt's router
	r.GET("/", ctl.Index)
	r.POST("/", ctl.IndexSubmission)
	r.Handler("GET", "/favicon.ico", http.NotFoundHandler())
	r.ServeFiles("/public/*filepath", http.Dir("./public"))
	// call appengine's main function
	appengine.Main()
}