package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.Handle("/nf", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
