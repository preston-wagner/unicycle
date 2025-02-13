package multithread

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestMappingMultithread(t *testing.T) {
	result := MappingMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.ToString)
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingMultithread() returned unexpected %s", result)
	}

	if len(MappingMultithread(nil, test_ext.ToString)) != 0 {
		t.Error("MappingMultithread(nil) should return a slice with length 0")
	}
}

func TestMappingMultithreadWithError(t *testing.T) {
	result, err := MappingMultithreadWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.ToStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingMultithreadWithError() returned unexpected %s", result)
	}

	result, err = MappingMultithreadWithError(nil, test_ext.ToStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingMultithreadWithError(nil) should return a slice with length 0")
	}

	_, err = MappingMultithreadWithError([]int{1, 2, 3, -1, 7, 8}, test_ext.ToStringErrIfNegative)
	if err == nil {
		t.Error("MappingMultithreadWithError should return error if any mapping functions do")
	}
}
