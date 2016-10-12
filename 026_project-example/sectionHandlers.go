package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func initSection(r *httprouter.Router) {
	r.GET("/section/something", some)
}

func some(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	d := something()
	serveTemplate(w, "som.gohtml", d)
}
