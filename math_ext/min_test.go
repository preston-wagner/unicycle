package math_ext

import (
	"testing"
)

func TestMin(t *testing.T) {
	if total := Min(1, 2, 3, 4, 5); total != 1 {
		t.Errorf("Min() returned wrong value %v", total)
	}
	if total := Min(1.0, 3.4, 5.6, 7.9, 9.1); total != 1.0 {
		t.Errorf("Min() returned wrong value %v", total)
	}
}
