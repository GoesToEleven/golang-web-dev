package main

import (
	"time"
	"fmt"
)

func main(){
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("01 02 2006 03:04"))
	fmt.Println(t.Format("01 02 2006 15:04"))
}
