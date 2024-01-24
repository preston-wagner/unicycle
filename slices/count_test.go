package slices

import (
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/test_ext"
)

func TestCount(t *testing.T) {
	result := Count([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.Odd)
	if !reflect.DeepEqual(result, 5) {
		t.Errorf("Count() returned unexpected %d", result)
	}

	if Count(nil, test_ext.Odd) != 0 {
		t.Error("Count(nil) should return 0")
	}
}

func TestCountWithError(t *testing.T) {
	result, err := CountWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, 5) {
		t.Errorf("CountWithError() returned unexpected %d", result)
	}

	result, err = CountWithError(nil, test_ext.OddErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if result != 0 {
		t.Error("CountWithError(nil) should return 0")
	}

	_, err = CountWithError([]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 0}, test_ext.OddErrIfNegative)
	if err == nil {
		t.Error("CountWithError should return an error if a test function did")
	}
}
