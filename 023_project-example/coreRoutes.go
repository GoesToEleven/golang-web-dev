package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	initSection(r)
	http.ListenAndServe(":8080", r)
}
