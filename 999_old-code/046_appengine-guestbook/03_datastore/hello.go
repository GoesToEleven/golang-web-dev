package guestbook

import (
	"html/template"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

// IMPORTANT
// you might need to run this at the terminal
// go get -u golang.org/x/net/context

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(res, greetings); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="POST">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

func sign(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	g := Greeting{
		Content: req.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/", http.StatusFound)
}
