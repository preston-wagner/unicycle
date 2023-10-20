package sets

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	set := make(Set[string])
	set.Add("lorem")
	set.Add("ipsum")
	if len(set) != 2 {
		t.Errorf("Error adding values to Set")
	}
	set.Add("ipsum")
	if len(set) != 2 {
		t.Errorf("Error adding redundant values to Set")
	}
	set.Remove("lorem")
	if len(set) != 1 {
		t.Errorf("Error removing values from Set")
	}
	set.Remove("lorem")
	if len(set.Values()) != 1 {
		t.Errorf("Error getting list of values from Set")
	}
	if !set.Has("ipsum") {
		t.Errorf("Removed value still in set")
	}
	if set.Has("lorem") {
		t.Errorf("Existing value not found in set")
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
