package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", user)
	http.HandleFunc("/logout", logOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}