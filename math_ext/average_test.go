package math_ext

import (
	"testing"
)

func TestAverage(t *testing.T) {
	if average := Average(1, 2, 3, 4, 5); average != 3 {
		t.Errorf("Average() returned wrong average %v", average)
	}
	if average := Average(1.0, 3.4, 5.6, 7.9, 9.1); average != 5.4 {
		t.Errorf("Average() returned wrong average %v", average)
	}
	if average := Average[int](); average != 0 {
		t.Errorf("Average() returned wrong average %v", average)
	}
}

func TestAverage64(t *testing.T) {
	if average := Average64(1, 2, 3, 4, 5); average != 3 {
		t.Errorf("Average64() returned wrong average %v", average)
	}
	if average := Average64(1.0, 3.4, 5.6, 7.9, 9.1); average != 5.4 {
		t.Errorf("Average64() returned wrong average %v", average)
	}
	if average := Average64[int](); average != 0.0 {
		t.Errorf("Average64() returned wrong average %v", average)
	}
}
