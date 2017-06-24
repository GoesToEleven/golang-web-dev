package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type SafeCounter struct {
	c int
	sync.Mutex
}

func (c *SafeCounter) Add() {
	c.Lock()
	defer c.Unlock()
	c.c++
}

var counter *SafeCounter = &SafeCounter{}

func main() {

	fmt.Println(&counter.c)

	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter.c)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		fmt.Println("1----", counter)
		fmt.Println("2----", x)
		x.Add()
		counter = x
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		fmt.Println(s, i, "Counter:", counter.c)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
