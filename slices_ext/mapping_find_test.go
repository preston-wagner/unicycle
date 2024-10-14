package slices_ext

import (
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestMappingFind(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result, ok := MappingFind(input, test_ext.ToStringIfOdd)
	if !ok {
		t.Errorf("MappingFind() returned unexpected ok=false")
	}
	if result != "1" {
		t.Errorf("MappingFind() returned unexpected result")
	}

	if _, ok := MappingFind(nil, test_ext.ToStringIfOdd); ok {
		t.Error("MappingFind(nil) should return a ok=false")
	}
}
