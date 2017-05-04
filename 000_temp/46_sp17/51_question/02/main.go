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
	m sync.Mutex
}

func (sc *SafeCounter) Add() {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.c++
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
		counter.Add()
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		//fmt.Println(s, i, "Counter:", counter.c) // causes a race; no lock here; multiple goroutines accessing to READ the value
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
