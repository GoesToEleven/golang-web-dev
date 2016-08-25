package mickeymouse

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/retrieve", noConfusion)
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	// set cookie
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id.String(),
		// Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(res, cookie)

	// set memcache
	ctx := appengine.NewContext(req)
	item1 := memcache.Item{
		Key:   id.String(),
		Value: []byte("McLeod"),
	}
	memcache.Set(ctx, &item1)

	fmt.Fprint(res, "EVERYTHING SET ID:"+id.String())
}

func noConfusion(res http.ResponseWriter, req *http.Request) {

	var html string

	// get cookie value
	var id string
	cookie, _ := req.Cookie("session-id")
	if cookie != nil {
		id = cookie.Value
		html += `
				<br>
				<p>Value from cookie: ` + id + `</p>
				`
	}

	// get memcache value
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, id)
	if item != nil {
		html += `
			<br>
			<p>
			Value from memcache: ` + string(item.Value) + `
			</p>
		`
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)
}
