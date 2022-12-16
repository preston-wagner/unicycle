package unicycle

import (
	"testing"
)

func TestKeys(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
	}
	result := Keys(input)
	if len(result) != len(input) {
		t.Errorf("Keys() failed to return the correct number of keys")
	}
	if (!Includes(result, "a")) || (!Includes(result, "b")) || (!Includes(result, "c")) {
		t.Errorf("Keys() failed to return expected values")
	}
}
