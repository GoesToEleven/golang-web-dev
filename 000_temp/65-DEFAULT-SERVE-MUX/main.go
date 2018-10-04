package main

import (
	"net/http"
	"log"
	"fmt"
	"html"
	"io"
)


type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from baz")
}

func main() {

	var bazhandler hotdog

	http.Handle("/baz", bazhandler)

	http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/bar/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})


	log.Fatal(http.ListenAndServe(":8080", nil))

}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from foo")
}


/*

type Handler
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

*/