package server

import (
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"net/http"
)

func protected(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := session.Get(req)
		u := db.DB[c.Value]
		if !u.Loggedin {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		h(w, req)
	}
}
