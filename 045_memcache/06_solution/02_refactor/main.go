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

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
	http.SetCookie(w, cookie)

	// set memcache
	ctx := appengine.NewContext(r)
	item1 := memcache.Item{
		Key:   id.String(),
		Value: []byte("McLeod"),
	}
	memcache.Set(ctx, &item1)

	fmt.Fprint(w, "EVERYTHING SET ID:"+id.String())
}

func noConfusion(w http.ResponseWriter, r *http.Request) {

	var html string

	// get cookie value
	var id string
	cookie, err := r.Cookie("session-id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cookie != nil {
		id = cookie.Value
		html += `
				<br>
				<p>Value from cookie: ` + id + `</p>
				`
	}

	// get memcache value
	ctx := appengine.NewContext(r)
	item, _ := memcache.Get(ctx, id)
	if item != nil {
		html += `
			<br>
			<p>
			Value from memcache: ` + string(item.Value) + `
			</p>
		`
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}
