package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/main.css", css)
	http.HandleFunc("/bali.jpg", bali)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="main.css">
</head>
<body>

<h1>HELLO WORLD!</h1>

</body>
</html>`)
}

func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.css")
}

func bali(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "bali.jpg")
}