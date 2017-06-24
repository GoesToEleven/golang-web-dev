package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SITE</title>
</head>
<body>
<h1>HOME</h1>
<h1><a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SITE</title>
</head>
<body>
<h1>ABOUT</h1>
<h1><a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SITE</title>
</head>
<body>

<h1>CONTACT</h1>
<h1><a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
</body>
</html>`)
}
