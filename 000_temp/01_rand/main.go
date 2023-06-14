package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 랜덤 함수 출력
	rand.Seed(time.Now().Unix())
	x := rand.Intn(3)
	fmt.Println(x)
}
