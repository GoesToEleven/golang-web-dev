package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Snoop int

func (h Snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path, "/")
	fmt.Println(fs[0])
	fmt.Println(fs[1])
}

func main() {
	var dog Snoop
	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", dog)

	http.ListenAndServe(":8080", mux)
}
