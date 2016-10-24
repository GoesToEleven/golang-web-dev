package main

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/032_sample-project/route"
)

func main() {

	primaryMux := http.NewServeMux()
	route.Set(primaryMux)
	route.Requests(primaryMux)

	// ListenAndServe is using default ServeMux
	// All requests that arrive are handled by the handler route.WrapRequest
	http.HandleFunc("/", route.WrapRequest)
	http.ListenAndServe(":8080", nil)
}