package calc

import "testing"

func TestSum(t *testing.T) {
	v := Sum(1, 2)
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
