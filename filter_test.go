package unicycle

import (
	"reflect"
	"testing"
)

func odd(input int) bool {
	return input%2 == 1
}

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, odd)
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("Filter() returned unexpected %d", result)
	}

	if len(Filter(nil, odd)) != 0 {
		t.Error("Filter(nil) should return a slice with length 0")
	}
}
