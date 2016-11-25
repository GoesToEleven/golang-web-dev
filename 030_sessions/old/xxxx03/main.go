package main

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
}

var dbUser map[string]user
var dbSession map[string]string
var tpl *template.Template

func init() {
	dbUser = make(map[string]user)
	dbSession = make(map[string]string)
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/login", login)
	http.HandleFunc("/user", usr)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

}

func usr(w http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(w, "user.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u := dbUser[un]
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Forbidden: passwords do not match", http.StatusForbidden)
			return
		}
		c := createSession(un)
		http.SetCookie(w, c)
		http.Redirect(w, req, "/user", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func getSession(req *http.Request) (string, error) {
	c, err := req.Cookie("session")
	if err != nil {
		return "", err
	}
	userName, ok := dbSession[c.Value]
	if !ok {
		return "", errors.New("No such session")
	}
	return userName, nil
}

func createSession(userName string) *http.Cookie {
	id := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		//Secure: true,
		HttpOnly: true,
	}
	dbSession[id.String()] = userName
	return c
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Error(w, "You were already logged out", http.StatusBadRequest)
		return
	}
	delete(dbSession, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")
		p2 := req.FormValue("password-confirm")
		if p != p2 {
			http.Error(w, "passwords do not match", http.StatusBadRequest)
			return
		}
		if u == "" {
			http.Error(w, "no username entered", http.StatusBadRequest)
			return
		}
		if p == "" {
			http.Error(w, "no password entered", http.StatusBadRequest)
			return
		}
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		dbUser[u] = user{u, bs}
		cookie := createSession(u)
		http.SetCookie(w, cookie)
		return
	}
	http.Redirect(w, req, "/user", http.StatusSeeOther)
}

// 1.
// signup
// username, password, confirm password
// check password on server
//
//
//
//
//
//
//
//
//
// s
//
