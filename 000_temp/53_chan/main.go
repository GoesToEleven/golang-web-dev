package main

import "fmt"

func foo(f, b chan int) {
	for i := 0; i < 10000; i++ {
		select {
		case v := <-f:
			fmt.Println("from foo:", v)
		case v := <-b:
			fmt.Println("from bar:", v)
		}
	}
	fmt.Println("about to exit")
}

func main() {
	f := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i < 5000; i++ {
			f <- i
		}
	}()
	go func() {
		for i := 0; i < 5000; i++ {
			b <- i
		}
	}()
	foo(f, b)
}
