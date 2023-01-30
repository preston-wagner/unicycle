package unicycle

import (
	"testing"
)

func TestSome(t *testing.T) {
	if !Some([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, odd) {
		t.Errorf("Some() returned false negative")
	}
	if !Some([]int{1, 3, 5, 7, 9}, odd) {
		t.Errorf("Some() returned false negative")
	}
	if Some([]int{2, 4, 6, 8, 0}, odd) {
		t.Errorf("Some() returned false positive")
	}
	if Some([]int{}, odd) {
		t.Errorf("Some() returned false positive")
	}
}
