package maps

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	result := Merge(
		map[string]int{
			"a": 1,
			"b": 2,
			"c": 4,
			"e": 9,
		},
		map[string]int{
			"c": 3,
			"d": 4,
			"e": 6,
			"f": 6,
		},
		map[string]int{
			"e": 5,
		},
	)
	if !reflect.DeepEqual(result, map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
	}) {
		t.Errorf("Merge() returned unexpected result")
	}
}
