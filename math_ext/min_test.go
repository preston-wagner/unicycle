package math_ext

import (
	"testing"
)

func TestMin(t *testing.T) {
	if min := Min(1, 2, 3, 4, 5); min != 1 {
		t.Errorf("Min() returned wrong value %v", min)
	}
	if min := Min(3.4, 5.6, 1.0, 7.9, 9.1); min != 1.0 {
		t.Errorf("Min() returned wrong value %v", min)
	}
	if min := Min[float32](); min != 0.0 {
		t.Errorf("Min() returned wrong value %v", min)
	}
}
