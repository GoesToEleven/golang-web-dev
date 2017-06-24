package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}

	var u user
	if e, ok := dbSessions[c.Value]; ok {
		u = dbUsers[e]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	e := dbSessions[c.Value]
	_, ok := dbUsers[e]
	return ok
}
