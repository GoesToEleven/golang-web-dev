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
	fmt.Println(fs[2])
}

func main() {
	var dog Snoop
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", dog)
	http.ListenAndServe(":8080", nil)
}
