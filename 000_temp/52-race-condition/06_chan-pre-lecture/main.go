package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	c := make(chan int)

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	go func() {
		counter += <-c
		fmt.Println("counter:\t", counter)
	}()

	for i := 0; i < gs; i++ {
		go func() {
			c <- 1
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	close(c)
	fmt.Println("total count:", counter)
}