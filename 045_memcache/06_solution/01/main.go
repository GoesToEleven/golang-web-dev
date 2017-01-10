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
	}
	http.SetCookie(w, cookie)

	// set memcache
	ctx := appengine.NewContext(r)
	item1 := memcache.Item{
		Key:   id.String(),
		Value: []byte("McLeod"),
	}
	err := memcache.Set(ctx, &item1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "EVERYTHING SET ID:"+id.String())
}

func noConfusion(w http.ResponseWriter, r *http.Request) {

	html := `
	<form method="POST">
	    <input type="text" name="koala">
	    <input type="submit" value="submit">
	</form>
	`

	if r.Method == "POST" {
		id := r.FormValue("koala")

		// get cookie value
		cookie, err := r.Cookie("session-id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if cookie != nil {
			html += `
			<br>
			<p>Value from cookie: ` + cookie.Value + `</p>
			`
		}

		// get memcache value
		ctx := appengine.NewContext(r)
		item, err := memcache.Get(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if item != nil {
			html += `
			<br>
			<p>
			Value from memcache: ` + string(item.Value) + `
			</p>
		`
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}
