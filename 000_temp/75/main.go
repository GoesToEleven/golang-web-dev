package main

import (
	"html/template"
	"net/http"
)

const s = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>CONTACT 4</title>
</head>
<body>

<h1>YOU ARE AT CONTACT 4</h1>

<a href="/">HOME</a>
<a href="/about">ABOUT</a>
<a href="/contact/2">CONTACT 2</a>
<a href="/contact/3">CONTACT 3</a>
<a href="/contact/4">CONTACT 4</a>



</body>
</html>`

var tpl *template.Template
var tpl2 *template.Template
var tpl3 *template.Template
var tpl4 *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	tpl2 = template.Must(template.New("pale").ParseFiles("templates/contact.gohtml"))

	tpl3 = template.Must(template.ParseFiles("templates/contact.gohtml"))

	tpl4 = template.Must(template.New("pale").Parse(s))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", bar)
	http.HandleFunc("/contact/2", con2)
	http.HandleFunc("/contact/3", con3)
	http.HandleFunc("/contact/4", con4)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func con2(w http.ResponseWriter, r *http.Request) {
	tpl2.ExecuteTemplate(w, "contact.gohtml", nil)
}


func con3(w http.ResponseWriter, r *http.Request) {
	tpl3.ExecuteTemplate(w, "contact.gohtml", nil)
}


func con4(w http.ResponseWriter, r *http.Request) {
	tpl4.ExecuteTemplate(w, "pale", nil)
}
