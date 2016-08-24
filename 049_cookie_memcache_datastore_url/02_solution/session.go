package dstore

import (
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {
	ctx := appengine.NewContext(req)

	var id, origin string
	var cookie *http.Cookie
	// try to get the id from the COOKIE
	origin = "COOKIE"
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		// try to get the id from the URL
		origin = "URL"
		id := req.FormValue("id")
		if id == "" {
			// no id, so create one BRAND NEW
			origin = "BRAND NEW"
			pid, err := uuid.NewV4()
			if err != nil {
				log.Errorf(ctx, "ERROR getID uuid.NewV4: %s", err)
				return id, err
			}
			id = pid.String()
		}
		// try storing id in a cookie for later use
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	log.Infof(ctx, "ID CAME FROM: %s", origin)
	return id, nil
}
