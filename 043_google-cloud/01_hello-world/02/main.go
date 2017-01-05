package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", index)
	appengine.Main()
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, World!")
}
