package server

import (
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"net/http"
)

func user(w http.ResponseWriter, req *http.Request) {
	c := session.Get(req)
	u := db.DB[c.Value]
	tpl.ExecuteTemplate(w, "user.gohtml", u)
}
