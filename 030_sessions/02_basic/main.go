package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"time"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"encoding/json"
)

type user struct {
	First, Last, UserName string
	Password []byte
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSession = map[string]string{}

func init() {
	dbUsersLoad()
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	defer dbUsersSave()
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/loggedin", loggedin)
	http.HandleFunc("/elapsed", timer) //fyi
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// signup
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u := req.FormValue("username")
		p1 := req.FormValue("password1")
		p2 := req.FormValue("password2")
		if dbUsers[u] {
			http.Error(w, "Username taken", http.StatusUnprocessableEntity)
			return
		}
		if p1 != p2 {
			http.Error(w, "Passwords do not match", http.StatusUnprocessableEntity)
			return
		}
		bs, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.MinCost)
		if err != nil {
			log.Fatalln("bcrypt didn't work,", err)
			return
		}
		dbUsers[u] = user{f,l,u,bs}
		createSession(w, req, u)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func createSession(w http.ResponseWriter, req *http.Request, u string) {
	sID := uuid.NewV4()
	c := &http.Cookie{
		Name: "session",
		Value: sID.String(),
		//Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	dbSession[c.Value] = u
	// u is username, the key for dbUsers
	// with the session ID which is stored as the cookie's value
	// we can get the username
	// which is the key for dbUsers
	// and we can now access all of the user information
}

func login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		// retrieve form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		// check credentials - Does the username exist in the db?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password incorrect", http.StatusForbidden)
			return
		}

		// check credentials - Does the password entered match the password for that user?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password incorrect", http.StatusForbidden)
			return
		}

		createSession(w, req, un)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	sID := getSession(w, req)
	fname := db[sID]
	if req.Method == http.MethodPost {
		fname = req.FormValue("firstname")
		db[sID] = fname
	}
	log.Printf("DB in index - %v\n\n", db) //fyi
	tpl.ExecuteTemplate(w, "index.gohtml", fname)
}



func access(w http.ResponseWriter, req *http.Request) {
	sID := getSession(w, req)
	fname := db[sID]                        // *** accessing session data based upon session ID !!!***
	log.Printf("DB in access - %v\n\n", db) //fyi
	tpl.ExecuteTemplate(w, "access.gohtml", fname)
}

func getSession(w http.ResponseWriter, req *http.Request) string {
	c, err := req.Cookie("session")
	log.Printf("cookie received from browser - %v", c) //fyi
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	timeLimitedSession(w, c)
	log.Printf("cookie returned to browser - %v", c) //fyi
	return c.Value
}

func timeLimitedSession(w http.ResponseWriter, c *http.Cookie) {
	// c.Secure = true
	c.HttpOnly = true
	c.MaxAge = 30
	http.SetCookie(w, c)
	startSessionTimer() //fyi
}

//fyi - time the life of the cookie
var sessionStartTime time.Time

//fyi
func startSessionTimer() {
	sessionStartTime = time.Now()
}

//fyi - this does not call getSession therefore the session's MaxAge is not reset
func timer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Session time elapsed in seconds %v", time.Now().Sub(sessionStartTime).Seconds()) //fyi
}

func dbUsersLoad() {
	f, err := os.Open("db/simulatedDB.json")
	if err != nil {
		log.Fatalln("error opening json file", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&dbUsers)
	if err != nil {
		log.Fatalln("error decoding json", err)
	}
}

func dbUsersSave() {
	f, err := os.Create("db/simulatedDB.json")
	if err != nil {
		log.Fatalln("error creating json file", err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(&dbUsers)
	if err != nil {
		log.Fatalln("error encoding json", err)
	}
}