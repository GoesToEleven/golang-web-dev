package main

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/server"
)

func main() {
	server.Configure()
	http.ListenAndServe(":8080", nil)
}