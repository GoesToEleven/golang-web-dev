package main

import (
	"github.com/GoesToEleven/golang-web-dev/031_sessions/02/mcookie"
	"net/http"
)

func user(w http.ResponseWriter, req *http.Request) {
	c := mcookie.GetCookie(req)
	cl := db[c.Value]
	tpl.ExecuteTemplate(w, "user.gohtml", cl)
}
