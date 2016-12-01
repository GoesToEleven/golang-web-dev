package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"time"
	"fmt"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	// set cookie fields
	// c.Secure = true // HTTPS only
	c.HttpOnly = true // no javascript access
	c.MaxAge = 20 	// 20 second time limit
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.uID]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	s.LastActivity = time.Now()
	dbSessions[c.Value] = s
	_, ok := dbUsers[s.uID]
	return ok
}

func notAuthorized(req *http.Request) bool {
	// cookie?
	c, err := req.Cookie("session")
	if err != nil {
		return true
	}
	// session?
	s, ok := dbSessions[c.Value]
	if !ok {
		return true
	}
	// user?
	if _, ok = dbUsers[s.uID]; !ok {
		return true
	}
	return false
}

func cleanSessions() {
	fmt.Println("CLEAN SESSIONS RUNNING")
	for k, v := range dbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 40) {
			delete(dbSessions, k)
		}
	}
}