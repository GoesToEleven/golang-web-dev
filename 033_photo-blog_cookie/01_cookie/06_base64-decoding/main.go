package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	data := foo()
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: data,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprintln(res, cookie.Value)

	bs := bar(cookie.Value)
	fmt.Fprintln(res, string(bs))
}

func bar(data string) []byte {
	bs, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		log.Println("Error decoding base64", err)
	}
	return bs
}

func foo() string {
	m := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	str := base64.URLEncoding.EncodeToString(bs)
	return str
}
