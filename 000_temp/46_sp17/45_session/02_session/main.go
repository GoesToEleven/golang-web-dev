package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID (email)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}

	var u user
	if uid, ok := dbSessions[c.Value]; ok {
		u = dbUsers[uid]
	}

	// process form submission
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		f := r.FormValue("first")
		l := r.FormValue("last")
		u = user{e, f, l}
		dbSessions[c.Value] = e
		dbUsers[e] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	e, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[e]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
