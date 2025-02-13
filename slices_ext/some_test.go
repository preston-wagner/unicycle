package slices_ext

import (
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestSome(t *testing.T) {
	if !Some([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd) {
		t.Errorf("Some() returned false negative")
	}
	if !Some([]int{1, 3, 5, 7, 9}, test_ext.Odd) {
		t.Errorf("Some() returned false negative")
	}
	if Some([]int{2, 4, 6, 8, 0}, test_ext.Odd) {
		t.Errorf("Some() returned false positive")
	}
	if Some([]int{}, test_ext.Odd) {
		t.Errorf("Some() returned false positive")
	}
}
