package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://169.24.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	io.WriteString(w, string(bs))
}




















