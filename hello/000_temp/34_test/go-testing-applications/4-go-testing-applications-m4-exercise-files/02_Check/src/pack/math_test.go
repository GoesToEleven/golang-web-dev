package pack

import (
	"math"
	"testing"
	"testing/quick"
)

func TestNewIntegrator(t *testing.T) {
	pi := PolyIntegrator{}
	ri := RiemannIntegrator{}
	var err float64

	f := func(min, max byte, coefs [3]byte) bool {
		c1 := float64(coefs[0]) / 10.
		c2 := float64(coefs[1]) / 10.
		c3 := float64(coefs[2]) / 10.
		r1 := pi.Integrate(float64(min), float64(max), c1, c2, c3)
		r2 := ri.Integrate(float64(min), float64(max), c1, c2, c3)
		if r1 == 0 {
			if r2 == 0 {
				return true
			} else {
				err = math.Abs((r1 - r2) / r2)
			}
		} else {
			err = math.Abs((r1 - r2) / r1)
		}
		if err < 0.01 {
			return true
		}

		return false
	}

	e := quick.Check(f, nil)
	if e != nil {
		t.Error(err)
		t.Error(e)
	}
}
