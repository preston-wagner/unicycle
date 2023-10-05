package slices

import (
	"reflect"
	"testing"
)

func TestSliceRefs(t *testing.T) {
	input := []string{"a", "b", "c"}

	inputPointers := SliceRefs(input)

	*inputPointers[1] = "d"

	if !reflect.DeepEqual(input, []string{"a", "d", "c"}) {
		t.Errorf("SliceRefs() failed to reference originals")
	}
}
