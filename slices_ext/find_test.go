package slices_ext

import (
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestFind(t *testing.T) {
	result, ok := Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd)
	if !ok {
		t.Error("Find() should have returned ok=true")
	}
	if result != 1 {
		t.Errorf("Find() returned unexpected %d", result)
	}

	_, ok = Find([]int{2, 4, 6, 8, 0}, test_ext.Odd)
	if ok {
		t.Error("Find() should have returned ok=false")
	}

	if _, ok := Find(nil, test_ext.Odd); ok {
		t.Error("Find(nil) should return a slice with length 0")
	}
}
