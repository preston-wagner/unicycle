package slices

import (
	"reflect"
	"testing"
)

type keyStruct struct {
	key   int
	value string
}

func keyStructToKey(value keyStruct) int {
	return value.key
}

func TestKeyBy(t *testing.T) {
	values := []keyStruct{
		{
			key:   1,
			value: "a",
		},
		{
			key:   2,
			value: "b",
		},
		{
			key:   3,
			value: "c",
		},
		{
			key:   3,
			value: "d",
		},
	}

	result := KeyBy(values, keyStructToKey)

	if !reflect.DeepEqual(result, map[int]keyStruct{
		1: {
			key:   1,
			value: "a",
		},
		2: {
			key:   2,
			value: "b",
		},
		3: {
			key:   3,
			value: "d",
		},
	}) {
		t.Error("KeyBy returned unexpected value", result)
	}
}
