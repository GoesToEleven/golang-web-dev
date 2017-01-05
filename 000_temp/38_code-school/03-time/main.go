package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	h := t.Hour()
	fmt.Println(h)
}
