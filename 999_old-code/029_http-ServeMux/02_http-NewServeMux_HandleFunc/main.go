package main

import (
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "doggy")
	})

	mux.HandleFunc("/cat/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "catty")
	})

	http.ListenAndServe(":8080", mux)
}
