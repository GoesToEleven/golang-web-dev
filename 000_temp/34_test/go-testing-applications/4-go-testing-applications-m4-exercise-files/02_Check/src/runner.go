package main

import (
	"fmt"
	"pack"
)

func main() {
	pi := pack.PolyIntegrator{}

	fmt.Println(pi.Integrate(0, 10, 3))
	fmt.Println(pi.Integrate(0, 10, 1, 0))

	fmt.Println(pi.Integrate(0, 10, -1, 5))

	ri := pack.RiemannIntegrator{}
	fmt.Println(ri.Integrate(0, 10, 3))
	fmt.Println(ri.Integrate(0, 10, 1, 0))

	fmt.Println(ri.Integrate(0, 10, -1, 5))
}
