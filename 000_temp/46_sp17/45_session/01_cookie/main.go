package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
	fmt.Fprint(w, cookie)
}
