package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from hotdogger")
}

func main() {
	var foo hotdog
	http.Handle("/", foo)
	http.ListenAndServe(":8080", nil)
}
