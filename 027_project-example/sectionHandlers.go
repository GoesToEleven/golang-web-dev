package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const pathSection = "/section/something"

func initSection(r *httprouter.Router) {
	r.GET(pathSection, somethingHandler)
}

func somethingHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	d := something()
	serveTemplate(w, "something.gohtml", d)
}
