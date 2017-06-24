package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="public/css/main.css">
</head>
<body>

<h1>HELLO WORLD!</h1>

</body>
</html>`)
}

func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/css/main.css")
}

func bali(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/img/bali.jpg")
}
