package multithread

import (
	"testing"

	"github.com/preston-wagner/unicycle/slices"
	"github.com/preston-wagner/unicycle/test_ext"
)

func TestMappingFindMultithread(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, ok := MappingFindMultithread(input, test_ext.ToStringIfOdd)
	if !ok {
		t.Errorf("MappingFindMultithread() returned unexpected ok=false")
	}
	result2 := slices.Mapping(slices.Filter(input, test_ext.Odd), test_ext.ToString)
	if !slices.Includes(result2, result) {
		t.Errorf("MappingFindMultithread() returned unexpected %s", result)
	}

	input = []int{2, 4, 6, 8, 0}
	_, ok = MappingFindMultithread(input, test_ext.ToStringIfOdd)
	if ok {
		t.Errorf("MappingFindMultithread() returned unexpected ok=true")
	}

	if _, ok := MappingFindMultithread(nil, test_ext.ToStringIfOdd); ok {
		t.Error("MappingFindMultithread(nil) should return ok=false")
	}
}
