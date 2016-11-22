package db

import "github.com/GoesToEleven/golang-web-dev/030_sessions/05/model"

var DB map[string]*model.User

func init() {
	DB = make(map[string]*model.User)
}
