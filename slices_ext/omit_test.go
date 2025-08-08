package slices_ext

import (
	"reflect"
	"testing"
)

func TestOmit(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := []int{1, 3, 5, 7, 9}
	result := Omit(slice1, slice2)
	if !reflect.DeepEqual(result, []int{2, 4, 6, 8}) {
		t.Errorf("Omit() returned unexpected %d", result)
	}

	if !reflect.DeepEqual(Omit(slice1, nil), slice1) {
		t.Error("Omit(slice1, nil) should return a slice with the same contents as the input")
	}

	if !reflect.DeepEqual(Omit(nil, slice2), []int{}) {
		t.Error("Omit(nil, *) should return an empty slice")
	}
}
