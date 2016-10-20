package mcookie

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"log"
)

// GetCookie will either
// (1) retrieve a cookie if it exists, then return it
// (2) create a cookie if it doesn't exist, then return it
func GetCookie(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}
	return cookie
}