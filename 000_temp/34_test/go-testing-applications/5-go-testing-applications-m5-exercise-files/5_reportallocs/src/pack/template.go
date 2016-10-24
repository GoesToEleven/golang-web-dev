package pack

import (
	"html/template"
	"os"
)

func SayHello(user User) {
	t, _ := template.New("").Parse("Hello {{.Name}}\n")
	t.Execute(os.Stdout, user)
}

type User struct {
	Name string
}

func NewUser(name string) User {
	result := User{name}
	return result
}
