package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	if total := Average([]int{1, 2, 3, 4, 5}); total != 3 {
		t.Errorf("Average() returned wrong total %v", total)
	}
	if total := Average([]float32{1.0, 3.4, 5.6, 7.9, 9.1}); total != 5.4 {
		t.Errorf("Average() returned wrong total %v", total)
	}
}
