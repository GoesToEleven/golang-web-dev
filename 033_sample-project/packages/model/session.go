package model

import "time"

// SessionData stores data for a session
type SessionData struct {
	User
	Expires time.Time
}