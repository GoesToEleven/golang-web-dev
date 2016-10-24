package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
}