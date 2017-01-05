package pack

import (
	"testing"
	"time"
)

func TestCanAddNumbers(t *testing.T) {
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
	result := Subtract(10, 5)

	if result != 5 {
		t.Error("Failed to subtract two numbers properly")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	if testing.Verbose() {
		t.Skip("Not implemented yet")
	}

}
