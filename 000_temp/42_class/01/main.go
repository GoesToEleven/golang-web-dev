package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	xs := req.Form["q"]
	for i, v := range xs {
		fmt.Fprintf(w, "Do my search %v - %v\n", i, v)
	}
}

// visit this page:
// http://localhost:8080/?q=dog
