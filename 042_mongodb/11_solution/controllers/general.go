package controllers

import (
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/11_solution/session"
	"html/template"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
