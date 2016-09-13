package main

import (
	"net/http"
)

type MyHandler int

func (m MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

func main() {
	var h MyHandler
	http.ListenAndServe(":8080", h)
}
