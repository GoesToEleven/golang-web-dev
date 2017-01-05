package pack

import (
	"math"
)

type Integrator interface {
	Integrate(min, max float64, coefs ...float64) float64
}

type PolyIntegrator struct{}

type XYFunc func(x float64) float64

func (pi *PolyIntegrator) Integrate(min, max float64, coefs ...float64) float64 {
	newCoefs := coefs[:]
	for i := range newCoefs {
		exp := float64(len(newCoefs) - i)
		newCoefs[i] = newCoefs[i] / exp
	}
	newCoefs = append(newCoefs, 0.)

	return calcPoly(newCoefs...)(max) - calcPoly(newCoefs...)(min)
}

func calcPoly(coefs ...float64) XYFunc {
	return func(x float64) float64 {
		result := 0.
		for i, coef := range coefs {
			result += coef * math.Pow(x, float64(len(coefs)-i-1))
		}
		return result
	}
}

type RiemannIntegrator struct{}

func (ri *RiemannIntegrator) IntegrateRiemann(min, max float64, f XYFunc) float64 {
	result := 0.

	numDivisions := 1000
	step := (max - min) / float64(numDivisions)

	for i := 0; i < numDivisions; i++ {
		xNow := min + step*float64(i)
		result += f(xNow)
	}

	return result * step
}

func (ri *RiemannIntegrator) Integrate(min, max float64, coefs ...float64) float64 {
	return ri.IntegrateRiemann(min, max, calcPoly(coefs...))
}
