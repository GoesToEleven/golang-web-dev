package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	val, ok := quick.Value(reflect.TypeOf(MyStruct{}),
		rand.New(rand.NewSource(time.Now().Unix())))

	if ok {
		fmt.Println(val.Interface().(MyStruct))
	}
}

type MyStruct struct {
	MyInt    int
	MyString string
	MySlice  []float32
}
