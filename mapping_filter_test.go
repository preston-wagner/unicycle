package unicycle

import (
	"fmt"
	"reflect"
	"testing"
)

func toStringIfOdd(input int) (string, bool) {
	if odd(input) {
		return fmt.Sprintf("%d", input), true
	}
	return "", false
}

func TestMappingFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	result := MappingFilter(input, toStringIfOdd)
	result2 := Mapping(Filter(input, odd), toString)
	if !reflect.DeepEqual(result, result2) {
		t.Errorf("MappingFilter() returned unexpected %s", result)
	}

	if len(MappingFilter(nil, toStringIfOdd)) != 0 {
		t.Error("MappingFilter(nil) should return a slice with length 0")
	}
}
