package main

import (
	"github.com/GoesToEleven/golang-web-dev/030_sessions/05/server"
	"net/http"
)

func main() {
	server.Configure()
	http.ListenAndServe(":8080", nil)
}
