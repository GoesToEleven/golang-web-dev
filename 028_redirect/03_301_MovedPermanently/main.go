package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
