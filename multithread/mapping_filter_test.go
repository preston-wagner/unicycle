package multithread

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/slices"
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

func TestMappingFilterMultithread(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := MappingFilterMultithread(input, toStringIfOdd)
	result2 := slices.Mapping(slices.Filter(input, odd), toString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterMultithread() returned unexpected %s", result)
	}

	if len(MappingFilterMultithread(nil, toStringIfOdd)) != 0 {
		t.Error("MappingFilterMultithread(nil) should return a slice with length 0")
	}
}

func TestMappingFilterMultithreadWithError(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, err := MappingFilterMultithreadWithError(input, toStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	result2 := slices.Mapping(slices.Filter(input, odd), toString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterMultithreadWithError() returned unexpected %s", result)
	}

	result, err = MappingFilterMultithreadWithError(nil, toStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingFilterMultithreadWithError(nil) should return a slice with length 0")
	}

	_, err = MappingFilterMultithreadWithError([]int{1, 2, 3, -1, 7, 8}, toStringIfOddErrIfNegative)
	if err == nil {
		t.Error("MappingFilterMultithreadWithError should return error if any mapping functions do")
	}
}
