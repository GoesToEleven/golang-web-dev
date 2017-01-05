package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	val, ok := quick.Value(reflect.TypeOf(MyInt(1)),
		rand.New(rand.NewSource(time.Now().Unix())))

	if ok {
		fmt.Println(val.Int())
	}
}

type MyInt int

func (mi MyInt) Generate(rand *rand.Rand, size int) reflect.Value {
	result := rand.Float32()*20. - 10.

	return reflect.ValueOf(int(result))
}
