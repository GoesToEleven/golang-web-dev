package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
