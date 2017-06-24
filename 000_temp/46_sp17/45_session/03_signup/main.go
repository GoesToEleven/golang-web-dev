package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password string
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
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		f := r.FormValue("first")
		l := r.FormValue("last")

		// is username taken, eg, user already signed up?
		if _, ok := dbUsers[e]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		id := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = e

		// store user in dbUsers
		u := user{e, p, f, l}
		dbUsers[e] = u

		// all done processing form submission
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
