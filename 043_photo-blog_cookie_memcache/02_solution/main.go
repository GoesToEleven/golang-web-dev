package mem

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/imgs/", fs)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("data")
		if err != nil {
			ctx := appengine.NewContext(req)
			log.Errorf(ctx, "error uploading photo: %s", err)
			// TODO: create error page to show user
		}
		cookie = uploadPhoto(src, hdr, cookie, req)
		http.SetCookie(res, cookie)
	}

	m := Model(cookie, req)
	tpl.ExecuteTemplate(res, "index.html", m)
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie := newVisitor(req)
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
}

func login(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" && req.FormValue("password") == "secret" {
		m := Model(cookie, req)
		m.State = true
		m.Name = req.FormValue("name")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(m, id, req)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}
