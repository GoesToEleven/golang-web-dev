package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	initCRM(r)
	http.ListenAndServe(":8080", r)
}