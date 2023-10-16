package multithread

import (
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestEveryMultithread(t *testing.T) {
	if EveryMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd, 10) {
		t.Errorf("EveryMultithread() returned false positive")
	}
	if !EveryMultithread([]int{1, 3, 5, 7, 9}, test_ext.Odd, 10) {
		t.Errorf("EveryMultithread() returned false negative")
	}
	if EveryMultithread([]int{2, 4, 6, 8, 0}, test_ext.Odd, 10) {
		t.Errorf("EveryMultithread() returned false positive")
	}
	if !EveryMultithread([]int{}, test_ext.Odd, 10) {
		t.Errorf("EveryMultithread() returned false negative")
	}
}
