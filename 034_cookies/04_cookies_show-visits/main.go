package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// disregard favicon requests
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		// try to get the cookie
		cookie, err := req.Cookie("my-cookie")
		// if there is no cookie, create one
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "my-cookie",
				Value: "0",
			}
		}
		// there is always a cookie at this point
		count, _ := strconv.Atoi(cookie.Value)
		count++
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)
	})
	http.ListenAndServe(":8080", nil)
}
