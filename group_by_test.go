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
		t.Error("KeyBy returned unexpected value", result)
	}
}
