package main

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"log"
	"html/template"
)

type user struct {
	Email string
	Fname string
	Loggedin bool
}

var tpl *template.Template
var db = make(map[string]user)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/user", bar)
	http.HandleFunc("/logout", logOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	c := getCookie(req)

	if req.Method == http.MethodPost {
		id := c.Value
		db[id] = user{
			Email: req.FormValue("email"),
			Fname: req.FormValue("fname"),
			Loggedin: req.FormValue("loggedin") == "on",
		}
		http.Redirect(w, req, "/user", http.StatusSeeOther)
	}

	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func bar(w http.ResponseWriter, req *http.Request) {
	c := getCookie(req)
	u := db[c.Value]
	if !u.Loggedin {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "user.gohtml", u)
}

func getCookie(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}
	return cookie
}

func logOut(w http.ResponseWriter, req *http.Request) {
		c := getCookie(req)
		u := db[c.Value]
		u.Loggedin = false
		db[c.Value] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
}