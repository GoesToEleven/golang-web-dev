package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Any code you want in this func")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}