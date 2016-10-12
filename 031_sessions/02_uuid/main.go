package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

}

// go get uuid
// https://github.com/nu7hatch/gouuid
// NewV4

// https://en.wikipedia.org/wiki/Universally_unique_identifier
