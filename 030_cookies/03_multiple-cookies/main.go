package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("fname")
	cookie2, err := req.Cookie("lname")
	fmt.Println(cookie.Value, err)
	fmt.Println(cookie2.Value, err)

	http.SetCookie(res, &http.Cookie{
		Name:  "fname",
		Value: "Todd",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "lname",
		Value: "McLeod",
	})
}
