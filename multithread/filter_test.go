package multithread

import (
	"errors"
	"reflect"
	"testing"
)

func TestFilterMultithread(t *testing.T) {
	result := FilterMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, odd)
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterMultithread() returned unexpected %d", result)
	}

	if len(FilterMultithread(nil, odd)) != 0 {
		t.Error("FilterMultithread(nil) should return a slice with length 0")
	}
}

func oddErrIfNegative(input int) (bool, error) {
	if input < 0 {
		return false, errors.New("negative number")
	}
	return input%2 == 1, nil
}

func TestFilterMultithreadWithError(t *testing.T) {
	result, err := FilterMultithreadWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, oddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterWithError() returned unexpected %d", result)
	}

	result, err = FilterMultithreadWithError(nil, oddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("FilterWithError(nil) should return a slice with length 0")
	}

	_, err = FilterMultithreadWithError([]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}, oddErrIfNegative)
	if err == nil {
		t.Error("FilterWithError should return an error if a filter function did")
	}
}
