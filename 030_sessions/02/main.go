package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	First    string
	LoggedIn bool
}

var db map[string]user
var tpl *template.Template

func init() {
	db = make(map[string]user)
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/usr", usr)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getSession(req)
	u := db[c.Value] // get user

	if req.Method == http.MethodPost {
		f := req.FormValue("first")
		in := req.FormValue("loggedin") == "on"
		u = user{f, in}
	}

	db[c.Value] = u // store user
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func usr(w http.ResponseWriter, req *http.Request) {
	c := getSession(req)
	u := db[c.Value]

	if !u.LoggedIn {
		fmt.Println("Redirecting to /")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "user.gohtml", u)
}

func getSession(req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
	}
	return c
}

func logout(w http.ResponseWriter, req *http.Request) {
	c := getSession(req)
	u := db[c.Value]   // retrieve user from db
	u.LoggedIn = false // change user value
	db[c.Value] = u    // store user in db
	c.MaxAge = -1      // expire cookie
	http.SetCookie(w, c)
	fmt.Println("\n\n***Redirecting to /usr")
	http.Redirect(w, req, "/usr", http.StatusSeeOther)
}
