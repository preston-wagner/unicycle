package maps

import (
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	original := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	inverted := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	if !reflect.DeepEqual(original, Invert(inverted)) {
		t.Errorf("Invert error")
	}
}
