package math_ext

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	result := Range(0, 6, 1)
	if !reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("Range() returned unexpected %v", result)
	}

	result2 := Range(3.0, 7.0, 0.3)
	if len(result2) != 14 {
		t.Errorf("Range() returned unexpected %v", result2)
	}
}
