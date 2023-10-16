package slices

import (
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestEvery(t *testing.T) {
	if Every([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd) {
		t.Errorf("Every() returned false positive")
	}
	if !Every([]int{1, 3, 5, 7, 9}, test_ext.Odd) {
		t.Errorf("Every() returned false negative")
	}
	if Every([]int{2, 4, 6, 8, 0}, test_ext.Odd) {
		t.Errorf("Every() returned false positive")
	}
	if !Every([]int{}, test_ext.Odd) {
		t.Errorf("Every() returned false negative")
	}
}
