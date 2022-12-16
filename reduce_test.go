package unicycle

import (
	"testing"
)

func sum(total, current int) int {
	return total + current
}

func TestReduce(t *testing.T) {
	result := Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, sum, 0)
	if result != 45 {
		t.Errorf("Reduce() returned unexpected %d", result)
	}
}
