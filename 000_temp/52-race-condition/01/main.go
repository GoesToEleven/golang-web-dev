package _1

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var mu sync.Mutex

func main() {

	const gs = 4
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// 			time.Sleep(time.Second * 2)
			runtime.Gosched() // tells runtime to yield processor to other goroutine
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
