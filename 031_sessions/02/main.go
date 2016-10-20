package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", user)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}