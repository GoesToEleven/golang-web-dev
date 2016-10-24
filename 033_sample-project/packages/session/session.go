package session

import "github.com/GoesToEleven/golang-web-dev/033_sample-project/packages/model"

// ActiveSessions stores users who have active sessions
var ActiveSessions map[int]model.SessionData

func init() {
	ActiveSessions = make(map[int]model.SessionData)
}