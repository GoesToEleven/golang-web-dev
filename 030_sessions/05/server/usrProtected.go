package server

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
)

func protected(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		c := session.Get(req)
		u := db.DB[c.Value]
		if !u.Loggedin {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		h(w, req)
	}
}
