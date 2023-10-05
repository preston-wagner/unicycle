package maps

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
	}
	result := Values(input)
	if len(result) != len(input) {
		t.Errorf("Values() failed to return the correct number of keys")
	}
	if !reflect.DeepEqual(result, []int{1, 1, 1}) {
		t.Errorf("Values() failed to return expected values")
	}
}
