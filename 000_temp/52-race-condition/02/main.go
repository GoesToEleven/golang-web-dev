package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	fmt.Println(runtime.NumCPU())

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine())
		fmt.Println(runtime.NumCPU())
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(counter)
}
