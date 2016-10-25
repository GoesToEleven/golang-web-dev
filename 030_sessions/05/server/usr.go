package server

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
)

func user(w http.ResponseWriter, req *http.Request) {
	c := session.Get(req)
	u := db.DB[c.Value]
	tpl.ExecuteTemplate(w, "user.gohtml", u)
}
