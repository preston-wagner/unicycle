package slices_ext

import (
	"reflect"
	"testing"
)

func keyStructToValue(value keyStruct) string {
	return value.value
}

func TestSortBy(t *testing.T) {
	values := []keyStruct{
		{
			key:   1,
			value: "a",
		},
		{
			key:   3,
			value: "d",
		},
		{
			key:   3,
			value: "c",
		},
		{
			key:   2,
			value: "b",
		},
	}

	sorted := SortBy(values, keyStructToKey)

	if !reflect.DeepEqual(sorted, []keyStruct{
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
			value: "d",
		},
		{
			key:   3,
			value: "c",
		},
	}) {
		t.Error("SortBy returned unexpected value", sorted)
	}

	sorted = SortBy(values, keyStructToValue)

	if !reflect.DeepEqual(sorted, []keyStruct{
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
	}) {
		t.Error("SortBy returned unexpected value", sorted)
	}
}
