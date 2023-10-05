package random

import (
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/sets"
)

func TestShuffle(t *testing.T) {
	original := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	shuffled := Shuffle(original)
	if reflect.DeepEqual(original, shuffled) {
		t.Errorf("Shuffle() failed to return a different order")
	}

	if !reflect.DeepEqual(sets.SetFromSlice(original), sets.SetFromSlice(shuffled)) {
		t.Error("Shuffle() failed to return the same elements")
	}
}
