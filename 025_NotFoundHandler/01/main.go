package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	// favicon.ico를 포함하지않으려면 다음 코드를 포함시켜줘야함
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
