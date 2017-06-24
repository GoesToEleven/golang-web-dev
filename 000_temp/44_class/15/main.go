package main

import (
	"fmt"
	"io"
	"net/http"
)

func init() {
	fmt.Println("hello from init")
}

func main() {
	fmt.Println("hello from main")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
