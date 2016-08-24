package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/imgs/", fs)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	fmt.Println("COOKIE VALUE", cookie.Value, "\n\n") // DEBUGGING SOLUTION HINT

	if tampered(cookie.Value) {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("data")
		if err != nil {
			log.Println("error uploading photo: ", err)
			// TODO: create error page to show user
		}
		cookie = uploadPhoto(src, hdr, cookie)
		http.SetCookie(res, cookie)
	}

	m := Model(cookie)
	tpl.ExecuteTemplate(res, "index.html", m)
}
