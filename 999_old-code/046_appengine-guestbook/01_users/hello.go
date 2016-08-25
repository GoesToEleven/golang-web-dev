package hello

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, req.URL.String())
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(res, "Hello, %v!", u)
}
