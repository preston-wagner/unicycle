package unicycle

import (
	"fmt"
	"reflect"
	"testing"
)

func toString(input int) string {
	return fmt.Sprintf("%d", input)
}

func TestMapping(t *testing.T) {
	result := Mapping([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, toString)
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("Mapping() returned unexpected %s", result)
	}

	if len(Mapping(nil, toString)) != 0 {
		t.Error("Mapping(nil) should return a slice with length 0")
	}
}
