package mem

import (
	"encoding/base64"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
	"strings"
)

func genCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = newVisitor(req)
		http.SetCookie(res, cookie)
		return cookie
	}

	// make sure set cookie uses our current structure
	if strings.Count(cookie.Value, "|") != 2 {
		cookie = newVisitor(req)
		http.SetCookie(res, cookie)
		return cookie
	}

	if tampered(cookie.Value) {
		cookie = newVisitor(req)
		http.SetCookie(res, cookie)
		return cookie
	}

	ctx := appengine.NewContext(req)
	id := strings.Split(cookie.Value, "|")[0]
	item, _ := memcache.Get(ctx, id)
	if item == nil {
		cookie = newVisitor(req)
		http.SetCookie(res, cookie)
		return cookie
	}

	return cookie
}

func makeCookie(mm []byte, id string, req *http.Request) *http.Cookie {

	// Anytime a cookie is created, let's print the id
	// The id is the key for the value in memcache
	// Having the id will allow us to lookup the value in memcache
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "ID: %s", id)

	b64 := base64.URLEncoding.EncodeToString(mm)

	// SEND DATA TO BE STORED IN MEMCACHE
	storeMemc([]byte(b64), id, req)

	// SEND DATA TO BE STORED IN COOKIE
	code := getCode(b64) // hmac
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}
