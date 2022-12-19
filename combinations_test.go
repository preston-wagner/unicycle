package unicycle

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	result := Combinations([]string{"a", "b", "c"})
	if !reflect.DeepEqual(result, [][]string{
		{"a", "b", "c"},
		{"a", "b"},
		{"a", "c"},
		{"a"},
		{"b", "c"},
		{"b"},
		{"c"},
		{},
	}) {
		t.Errorf("Concatenate() returned unexpected %s", result)
	}
	result = Combinations([]string{})
	if !reflect.DeepEqual(result, [][]string{
		{},
	}) {
		t.Errorf("Concatenate() returned unexpected %s", result)
	}
}
