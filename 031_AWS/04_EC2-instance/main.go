package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello from AWS.")
	if err != nil {
		fmt.Println(err)
	}
}