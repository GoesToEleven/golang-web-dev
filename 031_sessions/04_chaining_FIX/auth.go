package main

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/031_sessions/04_chaining_FIX/mcookie"
)

func Auth1(f func(w http.ResponseWriter, req *http.Request)) func(http.ResponseWriter, *http.Request) {
	c := mcookie.GetCookie(req)
	cl := db[c.Value]
	if !cl.Loggedin {
		return func(w, req){

		}
	}
}
