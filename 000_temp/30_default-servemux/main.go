package main

import (
	"fmt"
	"net/http"
)

func bigd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func DOGGGGGG")
}

func indx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func CATTTTTT")
}

func main() {
	http.Handle("/", indx)
	http.HandleFunc("/dog", bigd)
	http.ListenAndServe(":8080", nil)
}
