package unicycle

import (
	"reflect"
	"testing"
)

func TestPick(t *testing.T) {
	result := Pick(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}, []string{"b", "d", "e"})
	if !reflect.DeepEqual(result, map[string]int{
		"b": 2,
		"d": 4,
	}) {
		t.Errorf("Pick() returned wrong selection")
	}
}
