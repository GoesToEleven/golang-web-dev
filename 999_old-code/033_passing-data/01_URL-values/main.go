package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		val := req.URL.Query().Get(key)
		io.WriteString(res, "Do my search:"+val)
	})

	http.ListenAndServe(":8080", nil)
}

// visit this page:
// http://localhost:8080/?q="dog"
