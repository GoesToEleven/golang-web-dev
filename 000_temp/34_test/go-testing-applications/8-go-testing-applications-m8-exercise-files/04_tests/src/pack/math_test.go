package pack

import (
	"testing"
)

func TestPolyIntegrator_Integrate(t *testing.T) {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 3)

	if result != 30 {
		t.Fail()
	}
}

func TestPolyIntegrator_Integrate_line(t *testing.T) {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 1, 0)

	if result != 50 {
		t.Fail()
	}
}
