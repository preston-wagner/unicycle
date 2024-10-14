package slices_ext

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestMappingFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := MappingFilter(input, test_ext.ToStringIfOdd)
	result2 := Mapping(Filter(input, test_ext.Odd), test_ext.ToString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilter() returned unexpected %s", result)
	}

	if len(MappingFilter(nil, test_ext.ToStringIfOdd)) != 0 {
		t.Error("MappingFilter(nil) should return a slice with length 0")
	}
}

func TestMappingFilterWithError(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, err := MappingFilterWithError(input, test_ext.ToStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	result2 := Mapping(Filter(input, test_ext.Odd), test_ext.ToString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterWithError() returned unexpected %s", result)
	}

	result, err = MappingFilterWithError(nil, test_ext.ToStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingFilterWithError(nil) should return a slice with length 0")
	}

	_, err = MappingFilterWithError([]int{1, 2, 3, -1, 7, 8}, test_ext.ToStringIfOddErrIfNegative)
	if err == nil {
		t.Error("MappingFilterWithError should return error if any mapping functions do")
	}
}
