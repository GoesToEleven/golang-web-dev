package main

import (
	"fmt"
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	fmt.Println("route d")
	io.WriteString(res, "dog dog doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	fmt.Println("route c")
	io.WriteString(res, "cat cat cat")
}

// 다음 케이스에서 c와 d는 HandlerFunc이다.
func main() {
	// ex1
	// http.HandleFunc("/dog/", d)
	// http.HandleFunc("/cat/", c)

	// ex2
	http.Handle("/dog", http.HandleFunc(d))
	http.Handle("/cat", http.HandleFunc(c))

	http.ListenAndServe(":8080", nil)
}
