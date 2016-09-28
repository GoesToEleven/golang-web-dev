package main

import (
	"io"
	"net/http"
)

type myHandler int
func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.URL.RequestURI())
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
