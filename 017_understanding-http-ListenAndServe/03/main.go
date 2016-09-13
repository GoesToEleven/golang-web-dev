package main

import (
	"io"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, req.URL.RequestURI())
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
