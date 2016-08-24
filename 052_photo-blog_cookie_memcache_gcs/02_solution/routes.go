package pblog

import (
	"html/template"
	"net/http"
)

var tpls *template.Template

const gcsBucket = "learning-1130.appspot.com"

func init() {
	tpls = template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	// get session
	s := getSession(res, req)

	// upload photo
	if req.Method == "POST" {
		s.uploadPhoto()
	}

	err := tpls.ExecuteTemplate(res, "index.html", *s)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
