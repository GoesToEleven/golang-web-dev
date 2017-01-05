package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/candidates.png", pic)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/candidates.png">`)
}

func pic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "c.png")
}
