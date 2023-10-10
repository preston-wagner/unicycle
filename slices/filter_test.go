package slices

import (
	"errors"
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

func oddErrIfNegative(input int) (bool, error) {
	if input < 0 {
		return false, errors.New("negative number")
	}
	return input%2 == 1, nil
}

func TestFilterWithError(t *testing.T) {
	result, err := FilterWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, oddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
		t.Errorf("FilterWithError() returned unexpected %d", result)
	}

	result, err = FilterWithError(nil, oddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("FilterWithError(nil) should return a slice with length 0")
	}

	_, err = FilterWithError([]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}, oddErrIfNegative)
	if err == nil {
		t.Error("FilterWithError should return an error if a filter function did")
	}
}