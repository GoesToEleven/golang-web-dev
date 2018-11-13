package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", newplace)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/new", http.StatusSeeOther)
	io.WriteString(w, "You are at index")
}

func newplace(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are at newplace")
}
