package route

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	//http.HandleFunc("/user/", user.Show)
	//http.HandleFunc("/login", index)
	//http.HandleFunc("/logout", index)
	http.Handle("/public/", http.StripPrefix("../public", http.FileServer(http.Dir("/public"))))
}