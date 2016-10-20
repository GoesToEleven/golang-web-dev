package main

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/031_sessions/02/mcookie"
)

func logOut(w http.ResponseWriter, req *http.Request) {
	c := mcookie.GetCookie(req)
	cl := db[c.Value]
	cl.Loggedin = false
	db[c.Value] = cl
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
