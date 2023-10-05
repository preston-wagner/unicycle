package math

import (
	"testing"
)

func TestMax(t *testing.T) {
	if total := Max(1, 2, 3, 4, 5); total != 5 {
		t.Errorf("Max() returned wrong value %v", total)
	}
	if total := Max(1.0, 3.4, 5.6, 7.9, 9.1); total != 9.1 {
		t.Errorf("Max() returned wrong value %v", total)
	}
}
