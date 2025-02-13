package multithread

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/slices_ext"
	"github.com/nuvi/unicycle/test_ext"
)

func TestMappingFilterMultithread(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := MappingFilterMultithread(input, test_ext.ToStringIfOdd)
	result2 := slices_ext.Mapping(slices_ext.Filter(input, test_ext.Odd), test_ext.ToString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterMultithread() returned unexpected %s", result)
	}

	if len(MappingFilterMultithread(nil, test_ext.ToStringIfOdd)) != 0 {
		t.Error("MappingFilterMultithread(nil) should return a slice with length 0")
	}
}

func TestMappingFilterMultithreadWithError(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, err := MappingFilterMultithreadWithError(input, test_ext.ToStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	result2 := slices_ext.Mapping(slices_ext.Filter(input, test_ext.Odd), test_ext.ToString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilterMultithreadWithError() returned unexpected %s", result)
	}

	result, err = MappingFilterMultithreadWithError(nil, test_ext.ToStringIfOddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingFilterMultithreadWithError(nil) should return a slice with length 0")
	}

	_, err = MappingFilterMultithreadWithError([]int{1, 2, 3, -1, 7, 8}, test_ext.ToStringIfOddErrIfNegative)
	if err == nil {
		t.Error("MappingFilterMultithreadWithError should return error if any mapping functions do")
	}
}
