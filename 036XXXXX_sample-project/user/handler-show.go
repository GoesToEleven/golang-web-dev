package user

import (
	"fmt"
	"net/http"
)

func Show(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userID, isValid := Get(ctx)
	fmt.Fprintf(res, "User is: %+v, the ID is valid: %t\n", userID, isValid)
}

