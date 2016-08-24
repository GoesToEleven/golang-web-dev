package main

import (
	"io"
	"net/http"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}

func youUp(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty")
}

func main() {

	http.HandleFunc("/", upTown)
	http.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":8080", nil)
}
