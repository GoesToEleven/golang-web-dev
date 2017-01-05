package pack

import (
	"html/template"
	"os"
	"time"
)

func SayHello(user User) {
	t, _ := template.New("").Parse("Hello {{.Name}}\n")
	time.Sleep(1 * time.Millisecond)
	t.Execute(os.Stdout, user)
}

type User struct {
	Name string
}

func NewUser(name string) User {
	result := User{name}
	time.Sleep(3 * time.Millisecond)
	return result
}
