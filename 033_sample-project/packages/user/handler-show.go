package user

import (
	"fmt"
	"net/http"
)

func Show(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "User is: %+v", nil)
}