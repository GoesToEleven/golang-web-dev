package main

import (
	"fmt"
	"pack"
	"testing"
)

func main() {
	br := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pack.PrintWeather(90001)
		}
	})

	fmt.Println(br)
}
