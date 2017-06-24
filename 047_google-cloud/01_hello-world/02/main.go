package main

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	appengine.Main()
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
