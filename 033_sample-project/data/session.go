package model

import "time"

// SessionData stores data for a session
type SessionData struct {
	User
	Expires time.Time
}

// ActiveSessions stores users who have active sessions
var ActiveSessions map[int]SessionData

func init() {
	ActiveSessions = make(map[int]SessionData)
}
