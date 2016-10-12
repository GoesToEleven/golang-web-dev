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

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, welcome to bar. What can I get you?")
}

func foo(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", 303)

	// or his way
	// http.Redirect(w, req, "/", http.StatusSeeOther)

	// or this way
	//res.Header().Set("Location", "/")
	//res.WriteHeader(http.StatusSeeOther)
}
