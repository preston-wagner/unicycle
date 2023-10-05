package multithread

import (
	"testing"
)

func odd(input int) bool {
	return input%2 == 1
}

func TestEveryMultithread(t *testing.T) {
	if EveryMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, odd, 10) {
		t.Errorf("EveryMultithread() returned false positive")
	}
	if !EveryMultithread([]int{1, 3, 5, 7, 9}, odd, 10) {
		t.Errorf("EveryMultithread() returned false negative")
	}
	if EveryMultithread([]int{2, 4, 6, 8, 0}, odd, 10) {
		t.Errorf("EveryMultithread() returned false positive")
	}
	if !EveryMultithread([]int{}, odd, 10) {
		t.Errorf("EveryMultithread() returned false negative")
	}
}
