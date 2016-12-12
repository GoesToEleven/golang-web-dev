package main

import (
    "net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
        http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}