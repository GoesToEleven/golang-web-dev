package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const pathSection = "/section/something"

func initSection(r *httprouter.Router) {
	r.GET(pathSection, somethingHandler)
}

func somethingHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	d := something()
	serveTemplate(res, "something.gohtml", d)
}
