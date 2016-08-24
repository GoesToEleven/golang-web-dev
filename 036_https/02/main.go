package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)

	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", 301))

	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
