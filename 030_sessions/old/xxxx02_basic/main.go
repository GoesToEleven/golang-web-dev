package main

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type user struct {
	First, Last, UserName string
	Password              []byte
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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := user{}
	c, err := req.Cookie("session")
	if err == nil {
		un := dbSession[c.Value]
		u = dbUsers[un]
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// signup
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u := req.FormValue("username")
		p1 := req.FormValue("password1")
		p2 := req.FormValue("password2")
		if _, ok := dbUsers[u]; ok {
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
		dbUsers[u] = user{f, l, u, bs}
		createSession(w, req, u)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func createSession(w http.ResponseWriter, req *http.Request, u string) {
	sID := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
		//Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(w, c)

	// dbSession is a map
	// the key is the session ID
	// the session ID is stored in the cookie (it's the cookie's value)
	// the value of dbSession is the username
	// the username is the key for dbUsers
	// dbUsers is a map
	// the value in dbUsers is a user
	// with a session ID, we can access user information
	dbSession[c.Value] = u
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

func logout(w http.ResponseWriter, req *http.Request) {
	sID := getSession(w, req)
	delete(dbSession, sID)
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func getSession(w http.ResponseWriter, req *http.Request) string {
	c, err := req.Cookie("session")
	log.Printf("cookie received from browser - %v", c) //fyi
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return ""
	}
	log.Printf("cookie returned to browser - %v", c) //fyi
	return c.Value
}

func loggedin(w http.ResponseWriter, req *http.Request) {
	sID := getSession(w, req)
	u := dbUsers[sID] // *** accessing user data based upon session ID !!!***
	tpl.ExecuteTemplate(w, "loggedin.gohtml", u)
}

func dbUsersLoad() {
	f, err := os.Open("db/simulatedDB.json")
	if err != nil {
		log.Fatalln("error opening json file", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&dbUsers)
	if err != nil {
		log.Println("error decoding json", err)
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
