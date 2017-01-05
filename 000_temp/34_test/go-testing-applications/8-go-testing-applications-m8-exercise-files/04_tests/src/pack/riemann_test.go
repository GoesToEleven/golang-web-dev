package pack

import (
	"testing"
)

var ri RiemannIntegrator

func TestRiemannIntegrator_Integrate(t *testing.T) {
	result := ri.Integrate(0, 10, 3)

	if result != 30 {
		t.Fail()
	}
}
