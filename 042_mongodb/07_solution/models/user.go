package models

// changed Id type to string
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name" `
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
