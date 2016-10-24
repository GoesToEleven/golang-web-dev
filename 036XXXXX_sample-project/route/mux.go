package route

import (
	"context"
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/032_sample-project/user"
)

const uk = byte(0)

var pMux *http.ServeMux

// Set configures a mux to handle request routing
func Set(primaryMux *http.ServeMux) {
	pMux = primaryMux
}

func WrapRequest(w http.ResponseWriter, req *http.Request) {
	u := user.Get(w, req)
	ctx := req.Context()
	ctx = context.WithValue(ctx, uk, u)
	req = req.WithContext(ctx)
	pMux.ServeHTTP(w, req)
}