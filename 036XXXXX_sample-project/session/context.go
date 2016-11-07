package session

import (
	"context"
	"github.com/GoesToEleven/golang-web-dev/032_sample-project/user"
)

func Get(c context.Context, userKey byte) (user.User, bool) {
	// assertion that interface{} type returned by c.Value is type user.User
	val, ok := c.Value(userKey).(user.User)
	return val, ok
}
