package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
    <style>
    html {
    font-size: 48px;
    }
    </style>
</head>
<body>
Hello from index
<br><br>
<a href="/">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>about</title>
        <style>
    html {
    font-size: 48px;
    }
    </style>
</head>
<body>
Hello from about
<br><br>
<a href="/">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>contact</title>
        <style>
    html {
    font-size: 48px;
    }
    </style>
</head>
<body>
Hello from contact
<br><br>
<a href="/">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}
