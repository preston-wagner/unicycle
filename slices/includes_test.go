package slices

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	if !Includes([]int{6, 7, 8}, 7) {
		t.Errorf("Includes() returned a false negative")
	}
	if Includes([]string{"a", "b", "c"}, "d") {
		t.Errorf("Includes() returned a false positive")
	}
}
