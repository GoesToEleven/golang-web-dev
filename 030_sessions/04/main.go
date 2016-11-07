package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// ADDED chaining: protected(user)
	http.HandleFunc("/user", protected(user))
	http.HandleFunc("/logout", logOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := GetSession(req)

	if req.Method == http.MethodPost {
		id := c.Value
		db[id] = &User{
			First:    req.FormValue("fname"),
			Last:     req.FormValue("lname"),
			Loggedin: req.FormValue("loggedin") == "on",
		}
		http.Redirect(w, req, "/user", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func user(w http.ResponseWriter, req *http.Request) {
	c := GetSession(req)
	u := db[c.Value]
	tpl.ExecuteTemplate(w, "user.gohtml", u)
}

func logOut(w http.ResponseWriter, req *http.Request) {
	c := GetSession(req)
	u := db[c.Value]
	u.Loggedin = false
	db[c.Value] = u
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// ADDED:
func protected(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := GetSession(req)
		u := db[c.Value]
		if !u.Loggedin {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		h(w, req)
	}
}
