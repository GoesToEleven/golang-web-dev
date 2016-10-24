package pack

import (
	"testing"
)

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()
	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add numbers as expected")
		}
	}
}

func TestCanSubtractNumbers(t *testing.T) {
	t.Parallel()
	result := Subtract(10, 5)
	if result != 5 {
		t.Error("Failed to subtract two numbers properly")
	}
}
