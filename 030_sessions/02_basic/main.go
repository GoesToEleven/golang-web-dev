package main

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

var tpl *template.Template
var db = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/access", access)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(req)
	fname := db[c.Value]
	if req.Method == http.MethodPost {
		fname = req.FormValue("firstname")
		db[c.Value] = fname
	}
	http.SetCookie(w, c)
	measureCookieDuration()
	log.Println("DB in index - ", db, "Time elapsed in seconds", time.Now().Sub(startTime).Seconds()) // fyi
	tpl.ExecuteTemplate(w, "index.gohtml", fname)
}

func getCookie(req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			MaxAge: 15,
		}
	}
	return c
}

func access(w http.ResponseWriter, req *http.Request) {
	c := getCookie(req)
	fname := db[c.Value]
	log.Println("DB in access - ", db, "Time elapsed in seconds", time.Now().Sub(startTime).Seconds()) // fyi
	tpl.ExecuteTemplate(w, "access.gohtml", fname)
}

// for timing the life of the cookie
var startTime time.Time
func measureCookieDuration() {
	startTime = time.Now()
}