package slices

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestMapping(t *testing.T) {
	result := Mapping([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.ToString)
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("Mapping() returned unexpected %s", result)
	}

	if len(Mapping(nil, test_ext.ToString)) != 0 {
		t.Error("Mapping(nil) should return a slice with length 0")
	}
}

func TestMappingWithError(t *testing.T) {
	result, err := MappingWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.ToStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingWithError() returned unexpected %s", result)
	}

	result, err = MappingWithError(nil, test_ext.ToStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingWithError(nil) should return a slice with length 0")
	}

	_, err = MappingWithError([]int{1, 2, 3, -1, 7, 8}, test_ext.ToStringErrIfNegative)
	if err == nil {
		t.Error("MappingWithError should return error if any mapping functions do")
	}
}
