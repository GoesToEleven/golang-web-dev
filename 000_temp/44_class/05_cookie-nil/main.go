package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  "something",
		Value: "mcleod example",
	}
	http.SetCookie(w, c)
	io.WriteString(w, "cookie set")

}

func bar(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("something")
	if c == nil {
		fmt.Println("equals nil")
		// fmt.Println(c.Value) // invalid memory address or nil pointer dereference
	}
	fmt.Printf("%T\n%v", c, c.Value)
}
