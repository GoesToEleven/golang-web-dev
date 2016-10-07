package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var chs = make(chan string)

func main() {

	xs := []string{"a", "b", "c", "d", "e"}
	wg.Add(len(xs))

	for _, v := range xs {
		go count(v)
	}

	go func() {
		wg.Wait()
		close(chs)
	}()

	for v := range chs {
		fmt.Println(v)
	}

	fmt.Println("Did we ever get here?")
}

func count(a string) {
	for i := 0; i <= 100; i++ {
		chs <- fmt.Sprintln(a, "-", i)
	}
	wg.Done()
}
