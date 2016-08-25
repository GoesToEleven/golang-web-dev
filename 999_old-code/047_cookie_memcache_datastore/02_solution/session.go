package dstore

import (
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {

	var id string
	var cookie *http.Cookie
	cookie, err := req.Cookie("session-id")
	if err != nil {
		pid, err := uuid.NewV4()
		if err != nil {
			ctx := appengine.NewContext(req)
			log.Errorf(ctx, "ERROR getID uuid.NewV4: %s", err)
			return id, err
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: pid.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	return id, nil
}
