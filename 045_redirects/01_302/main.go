package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", canine)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	const url = "/dog"
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusFound)
	// or this way, a little more robust:
	// http.Redirect(res, req, url, http.StatusFound)
}

func canine(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "What's up dog.")
}
