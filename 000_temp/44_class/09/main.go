package main

import "fmt"

type user struct {
	first string
}

func main() {
	m := map[string]user{}
	t := m["nothing"]
	t.first = "mcleod"
	m["nothing"] = t
	fmt.Println(t)
	fmt.Println(m)
}
