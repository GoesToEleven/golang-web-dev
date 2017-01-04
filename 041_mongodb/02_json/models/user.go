package models

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}

// to learn about JSON field tags
// https://godoc.org/encoding/json#Marshal
