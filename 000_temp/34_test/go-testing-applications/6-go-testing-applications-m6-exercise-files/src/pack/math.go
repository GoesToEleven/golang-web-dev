// This package contains various tools for integrating functions.
package pack

import (
	"math"
)

type Integrator interface {
	Integrate(min, max float64, coefs ...float64) float64
}

type XYFunc func(x float64) float64

type PolyIntegrator struct{}

func (integrator *PolyIntegrator) Integrate(min, max float64, coefs ...float64) float64 {
	newCoefs := coefs[:]
	for i := range newCoefs {
		exp := float64(len(newCoefs) - i)
		newCoefs[i] = newCoefs[i] / exp
	}
	newCoefs = append(newCoefs, 0.)

	return CalcPoly(newCoefs...)(max) - CalcPoly(newCoefs...)(min)
}

func CalcPoly(coefs ...float64) XYFunc {
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
	return ri.IntegrateRiemann(min, max, CalcPoly(coefs...))
}

func QuickSort(nums ...float64) *[]float64 {
	partition := func(arr *[]float64, low, high int) int {
		a := *arr
		pivot := a[high]
		i := low
		for j := low; j < high; j++ {
			if a[j] <= pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[high] = a[high], a[i]
		return i
	}
	var sort func(*[]float64, int, int)
	sort = func(a *[]float64, low, high int) {
		if low < high {
			p := partition(a, low, high)
			sort(a, low, p-1)
			sort(a, p+1, high)
		}
	}
	a := nums[:]

	sort(&a, 0, len(nums)-1)

	return &a
}
