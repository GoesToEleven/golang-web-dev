package session

import (
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/11_solution/models"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

const Length int = 30

var Users = map[string]models.User{}       // user ID, user
var Sessions = map[string]models.Session{} // session ID, session
var LastCleaned time.Time = time.Now()

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	ck, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	ck.MaxAge = Length
	http.SetCookie(w, ck)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
		u = Users[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}
	_, ok = Users[s.UserName]
	// refresh session
	ck.MaxAge = Length
	http.SetCookie(w, ck)
	return ok
}

func Clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// for demonstration purposes
func Show() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
