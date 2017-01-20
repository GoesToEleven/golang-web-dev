package main

import (
	"net/http"
	"io"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Holy diggity dogger")
}

