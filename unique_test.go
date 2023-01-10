package unicycle

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	result := Unique([]int{1, 2, 3, 3, 1, 4, 5, 4})
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Unique() returned unexpected %d", result)
	}
}
