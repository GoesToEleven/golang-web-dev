package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		if req.URL.Path == "/favicon.ico" {
			http.NotFound(res, req)
			return
		}

		cookie, err := req.Cookie("my-cookie")
		fmt.Println(cookie, err)

		http.SetCookie(res, &http.Cookie{
			Name:  "my-cookie",
			Value: "some value",
		})
	})
	http.ListenAndServe(":8080", nil)
}
