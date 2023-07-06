package unicycle

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
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

	result := GroupBy(values, keyStructToKey)

	if !reflect.DeepEqual(result, map[int][]keyStruct{
		1: {{
			key:   1,
			value: "a",
		}},
		2: {{
			key:   2,
			value: "b",
		}},
		3: {
			{
				key:   3,
				value: "c",
			}, {
				key:   3,
				value: "d",
			},
		},
	}) {
		t.Error("GroupBy returned unexpected value", result)
	}
}

func TestGroupByConcurrently(t *testing.T) {
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

	result := GroupByConcurrently(values, keyStructToKey)

	if !reflect.DeepEqual(result[1], []keyStruct{{
		key:   1,
		value: "a",
	}}) {
		t.Error("GroupByConcurrently returned unexpected value", result)
	}

	if !reflect.DeepEqual(result[2], []keyStruct{{
		key:   2,
		value: "b",
	}}) {
		t.Error("GroupByConcurrently returned unexpected value", result)
	}

	if !reflect.DeepEqual(SetFromSlice(result[3]), SetFromSlice([]keyStruct{
		{
			key:   3,
			value: "c",
		},
		{
			key:   3,
			value: "d",
		},
	})) {
		t.Error("GroupByConcurrently returned unexpected value", result)
	}
}
