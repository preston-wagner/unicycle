package slices_ext

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd)
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("Filter() returned unexpected %d", result)
	}

	if len(Filter(nil, test_ext.Odd)) != 0 {
		t.Error("Filter(nil) should return a slice with length 0")
	}
}

func TestFilterWithError(t *testing.T) {
	result, err := FilterWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterWithError() returned unexpected %d", result)
	}

	result, err = FilterWithError(nil, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("FilterWithError(nil) should return a slice with length 0")
	}

	_, err = FilterWithError([]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err == nil {
		t.Error("FilterWithError should return an error if a filter function did")
	}
}
