package main

import "net/http"

func loadRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./files"))))
}
