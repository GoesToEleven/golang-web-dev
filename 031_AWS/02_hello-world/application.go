package main

import (
    "net/http"
    "os"
	"io"
)

func main() {
    port := os.Getenv("PORT")
        if port == "" {
            port = "5000"
        }

	http.HandleFunc("/", index)
        http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, check out the size of my server, I'm running on AWS.")
}