package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	Id       string
	First    string
	Last     string
	Email    string
	Password string
}

var tpl *template.Template

var muser = map[string]user{}      // key is userid, value is user
var msession = map[string]string{} // key is sessionid, value is userid

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/other", other)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mysess")
	if err != nil {
		u := uuid.NewV4()
		c = &http.Cookie{
			Name:     "mysess",
			Value:    u.String(),
			Path:     "/",
			HttpOnly: true,
		}
	}
	uid := msession[c.Value]
	usr := muser[uid]

	fmt.Println("SESSION ID", c.Value)

	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", usr)
}

func other(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mysess")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uid := msession[c.Value]
	usr := muser[uid]

	fmt.Println("SESSION ID", c.Value)

	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "other.gohtml", usr)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mysess")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	delete(msession, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	//fmt.Fprint(w, "You've been logged out")
}

func login(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mysess")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		em := r.FormValue("email")
		pw := r.FormValue("password")
		// get all of the usernames and passwords
		m := map[string]user{}
		for _, v := range muser {
			m[v.Email] = v
		}
		usr, ok := m[em]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusForbidden)
			return
		}
		if usr.Password != pw {
			http.Redirect(w, r, "/login", http.StatusForbidden)
			return
		}
		msession[c.Value] = usr.Id
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mysess")

	// if no cookie, redirect to index
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// user has already signed up, so redirect
	uid := msession[c.Value]
	usr := muser[uid]
	if usr.Email != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		fn := r.FormValue("first")
		ln := r.FormValue("last")
		em := r.FormValue("email")
		pw := r.FormValue("password")
		u := uuid.NewV4()
		usr = user{
			Id:       u.String(),
			First:    fn,
			Last:     ln,
			Email:    em,
			Password: pw,
		}
		muser[usr.Id] = usr
		msession[c.Value] = usr.Id
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
