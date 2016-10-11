package main

import (
	"net/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
