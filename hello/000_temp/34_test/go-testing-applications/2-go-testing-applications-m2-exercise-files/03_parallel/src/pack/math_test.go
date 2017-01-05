package pack

import (
	"testing"
	"time"
)

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping long tests")
	}
	result := Add(1, 2)
	time.Sleep(1 * time.Second)
	if result != 3 {
		t.Error("Failed to add one and two")
	}

	result = Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("Failed to add more than two numbers")
	}
}

func TestCanSubtractNumbers(t *testing.T) {
	t.Parallel()
	result := Subtract(10, 5)
	time.Sleep(1 * time.Second)
	if result != 5 {
		t.Error("Failed to subtract two numbers properly")
	}
}
