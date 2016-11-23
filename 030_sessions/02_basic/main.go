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
	log.Printf("DB in index - %v - Time elapsed in seconds %v\n\n", db, time.Now().Sub(startTime).Seconds()) // fyi
	tpl.ExecuteTemplate(w, "index.gohtml", fname)
}

func getCookie(req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	log.Printf("cookie received back from browser - %v", c) //fyi
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}
	c.MaxAge = 30
	log.Printf("cookie returned by getCookie - %v", c) // fyi
	return c
}

func access(w http.ResponseWriter, req *http.Request) {
	c := getCookie(req)
	fname := db[c.Value]
	log.Printf("DB in access - %v - Time elapsed in seconds %v\n\n", db, time.Now().Sub(startTime).Seconds()) // fyi
	tpl.ExecuteTemplate(w, "access.gohtml", fname)
}

// for timing the life of the cookie
var startTime time.Time
func measureCookieDuration() {
	startTime = time.Now()
}