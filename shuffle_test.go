package unicycle

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	original := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	shuffled := Shuffle(original)
	if reflect.DeepEqual(original, shuffled) {
		t.Errorf("Shuffle() failed to return a different order")
	}

	if !reflect.DeepEqual(SetFromSlice(original), SetFromSlice(shuffled)) {
		t.Error("Shuffle() failed to return the same elements")
	}
}
