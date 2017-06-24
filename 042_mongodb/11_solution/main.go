package main

import (
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/11_solution/controllers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.SignUp)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
