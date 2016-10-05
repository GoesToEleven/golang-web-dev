package main

import (
	"net/http"
	"fmt"
)

type myHandler int
func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, req.URL.RequestURI())
	fmt.Fprintln(res, req.URL.Path)
	fmt.Fprintln(res, req.URL)
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
