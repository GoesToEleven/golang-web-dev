package main

import (
	"io"
	"net/http"
)

type DogHandler int
func (d DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type CatHandler int
func (c CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func main() {
	var dog DogHandler
	var cat CatHandler
	http.Handle("/", cat)
	http.Handle("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}