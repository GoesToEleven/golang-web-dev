package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req", req)
	fmt.Println("req.URL", req.URL)
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
