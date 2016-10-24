package main

import (
	"fmt"
	"math/rand"
	"pack"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	//quick.Value
	val, _ := quick.Value(reflect.TypeOf(1),
		rand.New(rand.NewSource(time.Now().Unix())))
	fmt.Println(val.Int())

	//quick.Check
	integrator := pack.PolyIntegrator{}
	fmt.Println(integrator.Integrate(0, 10, 3))
	fmt.Println(integrator.Integrate(0, 10, 1, 0))

	ri := pack.RiemannIntegrator{}

	fmt.Println(ri.Integrate(0, 10, 3))

	fmt.Println(ri.Integrate(0, 10, 1, 0))

	fmt.Println(pack.QuickSort(5, 3, 8, 1, 65.5))
}
