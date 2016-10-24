package user

import (
	"context"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

// These are not exported so it is impossible for things outside
// of this package to access the data in the context
type userkey byte

const uk = userkey(0)

var hand http.Handler

func GetUserWrapper(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// Get user from request, accessing the cookies and
	// looking up the actual user data in the database
	u := User{
		Name: "Bob",
		Age:  42,
	}
	ctx = context.WithValue(ctx, uk, u)
	req = req.WithContext(ctx)
	hand.ServeHTTP(res, req)
}

func GetUser(ctx context.Context) (User, bool) {
	val, ok := ctx.Value(uk).(User)
	return val, ok
}

func SetHandler(handler http.Handler) {
	hand = handler
}
