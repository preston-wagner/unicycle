package unicycle

import (
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
