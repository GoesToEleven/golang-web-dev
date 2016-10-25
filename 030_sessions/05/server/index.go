package server

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/session"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/db"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/model"
)

func index(w http.ResponseWriter, req *http.Request) {
	c := session.Get(req)

	if req.Method == http.MethodPost {
		id := c.Value
		db.DB[id] = &model.User{
			First: req.FormValue("fname"),
			Last: req.FormValue("lname"),
			Loggedin: req.FormValue("loggedin") == "on",
		}
		http.Redirect(w, req, "/user", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}
