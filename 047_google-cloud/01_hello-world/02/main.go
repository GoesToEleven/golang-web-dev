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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
