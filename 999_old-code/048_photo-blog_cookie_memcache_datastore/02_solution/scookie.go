package mem

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func getCookie(res http.ResponseWriter, req *http.Request) (*http.Cookie, error) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie, err = newVisitor(req)
		if err != nil {
			ctx := appengine.NewContext(req)
			log.Errorf(ctx, "ERROR getCookie newVisitor: %s", err)
			return nil, err
		}
		http.SetCookie(res, cookie)
		return cookie, nil
	}

	return cookie, nil
}

func makeCookie(m model, id string, req *http.Request) (*http.Cookie, error) {
	ctx := appengine.NewContext(req)

	// DATASTORE
	err := storeDstore(m, id, req)
	if err != nil {
		log.Errorf(ctx, "ERROR makeCookie storeDstore: %s", err)
		return nil, err
	}

	// MEMCACHE
	err = storeMemc(m, id, req)
	if err != nil {
		log.Errorf(ctx, "ERROR makeCookie storeMemc: %s", err)
		return nil, err
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}
