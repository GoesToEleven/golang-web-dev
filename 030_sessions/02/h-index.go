package main

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/031_sessions/02/mcookie"
)

func index(w http.ResponseWriter, req *http.Request) {
	c := mcookie.GetCookie(req)

	if req.Method == http.MethodPost {
		id := c.Value
		db[id] = client{
			Email: req.FormValue("email"),
			Fname: req.FormValue("fname"),
			Loggedin: req.FormValue("loggedin") == "on",
		}
		http.Redirect(w, req, "/user", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}
