package main

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

// GetSession will either (1) retrieve a cookie if it exists, then return it or (2) create a cookie if it doesn't exist, then return it.
func GetSession(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}
	return cookie
}
