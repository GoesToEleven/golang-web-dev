package main

import (
	"net/http"
	"io"
)


type hotdog int

// func receiver identifier(params) return(s) {}
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "WASSSSSUP!!!!!")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

/*

type Handler
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

*/