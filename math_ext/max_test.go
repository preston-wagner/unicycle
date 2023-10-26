package math_ext

import (
	"testing"
)

func TestMax(t *testing.T) {
	if max := Max(1, 2, 3, 4, 5); max != 5 {
		t.Errorf("Max() returned wrong value %v", max)
	}
	if max := Max(1.0, 3.4, 5.6, 7.9, 9.1); max != 9.1 {
		t.Errorf("Max() returned wrong value %v", max)
	}
	if max := Max[float32](); max != 0.0 {
		t.Errorf("Max() returned wrong value %v", max)
	}
}
