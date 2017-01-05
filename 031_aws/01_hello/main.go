package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}
