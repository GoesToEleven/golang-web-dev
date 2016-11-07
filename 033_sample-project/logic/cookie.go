package logic

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

func Get(w http.ResponseWriter, req *http.Request) *http.Cookie {
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
		http.SetCookie(w, cookie)
	}
	return cookie
}
