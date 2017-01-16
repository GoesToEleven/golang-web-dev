package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello from a docker container")
}
