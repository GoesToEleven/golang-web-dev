package pack

import (
	"fmt"
)

const FOO = iota

func Example() {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 3)

	fmt.Println(result)

	// Output:
	// 30
}
