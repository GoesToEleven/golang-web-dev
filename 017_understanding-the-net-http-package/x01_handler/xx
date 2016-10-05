package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, req.URL.RequestURI())
	fmt.Fprintln(res, req.URL.Path)
	fmt.Fprintln(res, req.URL)}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}