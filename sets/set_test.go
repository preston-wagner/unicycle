package sets

import (
	"reflect"
	"testing"
)

var slice = []string{
	"lorem",
	"ipsum",
}

func TestSet(t *testing.T) {
	set := Set[string]{}
	set.Add(slice...)
	if len(set) != len(slice) {
		t.Errorf("Error adding values to Set")
	}
	set.Add(slice[1])
	if len(set) != 2 {
		t.Errorf("Error adding redundant values to Set")
	}
	set.Remove(slice[0])
	if len(set) != 1 {
		t.Errorf("Error removing values from Set")
	}
	set.Remove(slice[0])
	if len(set.Values()) != 1 {
		t.Errorf("Error getting list of values from Set")
	}
	if set.Has(slice[0]) {
		t.Errorf("Removed value still in set")
	}
	if !set.Has(slice[1]) {
		t.Errorf("Existing value not found in set")
	}
}

func TestSetFromSlice(t *testing.T) {
	set := SetFromSlice(slice)
	if len(set) != len(slice) {
		t.Errorf("Error adding via SetFromSlice, unexpected length")
	}
	for _, value := range slice {
		if !set.Has(value) {
			t.Errorf("Existing value not found in set")
		}
	}
	for _, value := range slice {
		set.Add(value)
	}
	if len(set) != len(slice) {
		t.Errorf("Error adding duplciate values, unexpected length")
	}
}

func TestUnion(t *testing.T) {
	union := Union(SetFromSlice([]int{1, 2, 3}), SetFromSlice([]int{3, 4, 5}), SetFromSlice([]int{5, 6, 7}))
	if !reflect.DeepEqual(union, SetFromSlice([]int{1, 2, 3, 4, 5, 6, 7})) {
		t.Error("Union failed", union)
	}

	union = Union[int]()
	if !reflect.DeepEqual(union, SetFromSlice([]int{})) {
		t.Error("Union with no sets should return an empty set")
	}

	union = Union(SetFromSlice([]int{1, 2, 3, 4}))
	if !reflect.DeepEqual(union, SetFromSlice([]int{1, 2, 3, 4})) {
		t.Error("Union with a single set should return a copy of that set")
	}
}

func TestIntersection(t *testing.T) {
	intersection := Intersection(SetFromSlice([]int{1, 2, 3, 4}), SetFromSlice([]int{2, 3, 4, 5}), SetFromSlice([]int{3, 4, 5, 6}))
	if !reflect.DeepEqual(intersection, SetFromSlice([]int{3, 4})) {
		t.Error("Intersection failed", intersection)
	}

	intersection = Intersection[int]()
	if !reflect.DeepEqual(intersection, SetFromSlice([]int{})) {
		t.Error("Intersection with no sets should return an empty set")
	}

	intersection = Intersection(SetFromSlice([]int{1, 2, 3, 4}))
	if !reflect.DeepEqual(intersection, SetFromSlice([]int{1, 2, 3, 4})) {
		t.Error("Intersection with a single set should return a copy of that set")
	}
}

func TestDifference(t *testing.T) {
	difference := SetFromSlice([]int{1, 2, 3, 4}).Difference(SetFromSlice([]int{3, 4, 5, 6}), SetFromSlice([]int{5, 6, 7, 8}))
	if !reflect.DeepEqual(difference, SetFromSlice([]int{1, 2})) {
		t.Error("Difference failed", difference)
	}
}
