package multithread

import (
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestFilterMultithread(t *testing.T) {
	result := FilterMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd)
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterMultithread() returned unexpected %d", result)
	}

	if len(FilterMultithread(nil, test_ext.Odd)) != 0 {
		t.Error("FilterMultithread(nil) should return a slice with length 0")
	}
}

func TestFilterMultithreadWithError(t *testing.T) {
	result, err := FilterMultithreadWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterWithError() returned unexpected %d", result)
	}

	result, err = FilterMultithreadWithError(nil, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("FilterWithError(nil) should return a slice with length 0")
	}

	_, err = FilterMultithreadWithError([]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err == nil {
		t.Error("FilterWithError should return an error if a filter function did")
	}
}
