package user

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/032_sample-project/cookie"
)

func Get(w http.ResponseWriter, req *http.Request) User {
	cookie := cookie.Get(w, req)
	return User{
		Name: "Bob",
		Age:  42,
	}
}
