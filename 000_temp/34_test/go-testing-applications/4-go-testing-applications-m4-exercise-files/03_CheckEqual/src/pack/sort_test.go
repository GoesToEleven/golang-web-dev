package pack

import (
	"sort"
	"testing"
	"testing/quick"
)

func TestQuickSort(t *testing.T) {
	err := quick.CheckEqual(quickSortWrapper, sortWrapper, nil)

	if err != nil {
		t.Error(err)
	}
}

func quickSortWrapper(vals []float64) *[]float64 {
	return QuickSort(vals...)
}

func sortWrapper(vals []float64) *[]float64 {
	v := sort.Float64Slice(vals)
	v.Sort()

	f64 := []float64(v)

	return &f64
}
