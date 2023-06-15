package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	dog := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a dog!\n")
	}

	me := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a yongari #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
