package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, World!")
}
