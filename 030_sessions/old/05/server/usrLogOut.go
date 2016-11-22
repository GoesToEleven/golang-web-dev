package server

import (
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"net/http"
)

func logOut(w http.ResponseWriter, req *http.Request) {
	c := session.Get(req)
	u := db.DB[c.Value]
	u.Loggedin = false
	db.DB[c.Value] = u
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
