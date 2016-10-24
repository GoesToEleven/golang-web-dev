package route

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/032_sample-project/user"
)

func Requests(mux *http.ServeMux) {
	mux.HandleFunc("/", user.Show)
	mux.HandleFunc("/login", logIn)
	mux.HandleFunc("/logout", logOut)
	mux.HandleFunc("/user/", user)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
}


