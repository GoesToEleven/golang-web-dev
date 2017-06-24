package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Holy diggity dogger")
}
