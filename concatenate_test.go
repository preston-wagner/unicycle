package unicycle

import (
	"reflect"
	"testing"
)

func TestConcatenate(t *testing.T) {
	result := Concatenate([]string{"a", "b"}, []string{"c", "d", "e", "f", "g", "h"}, []string{"i"})
	if !reflect.DeepEqual(result, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}) {
		t.Errorf("Concatenate() returned unexpected %s", result)
	}
}
