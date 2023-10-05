package slices

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func toStringIfOdd(input int) (string, bool) {
	if odd(input) {
		return fmt.Sprintf("%d", input), true
	}
	return "", false
}

func toStringIfOddErrIfNegative(input int) (string, bool, error) {
	if input < 0 {
		return "", false, errors.New("toStringIfOddErrIfNegative(): negative number")
	}
	if odd(input) {
		return fmt.Sprintf("%d", input), true, nil
	}
	return "", false, nil
}

func TestMappingFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := MappingFilter(input, toStringIfOdd)
	result2 := Mapping(Filter(input, odd), toString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilter() returned unexpected %s", result)
	}

	if len(MappingFilter(nil, toStringIfOdd)) != 0 {
		t.Error("MappingFilter(nil) should return a slice with length 0")
	}
}

func TestMappingFilterWithError(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, err := MappingFilterWithError(input, toStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	result2 := Mapping(Filter(input, odd), toString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterWithError() returned unexpected %s", result)
	}

	result, err = MappingFilterWithError(nil, toStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingFilterWithError(nil) should return a slice with length 0")
	}

	_, err = MappingFilterWithError([]int{1, 2, 3, -1, 7, 8}, toStringIfOddErrIfNegative)
	if err == nil {
		t.Error("MappingFilterWithError should return error if any mapping functions do")
	}
}
