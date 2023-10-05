package multithread

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nuvi/unicycle/sets"
)

type keyStruct struct {
	key   int
	value string
}

func keyStructToKey(value keyStruct) int {
	return value.key
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

	if !reflect.DeepEqual(sets.SetFromSlice(result[3]), sets.SetFromSlice([]keyStruct{
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

func keyStructToKeyWithError(value keyStruct) (int, error) {
	if value.key%2 == 1 {
		return 0, errors.New("odd")
	}
	return value.key, nil
}

func TestGroupByConcurrentlyWithError(t *testing.T) {
	values := []keyStruct{
		{
			key:   2,
			value: "a",
		},
		{
			key:   4,
			value: "b",
		},
		{
			key:   6,
			value: "c",
		},
		{
			key:   6,
			value: "d",
		},
	}

	result, err := GroupByConcurrentlyWithError(values, keyStructToKeyWithError)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result[2], []keyStruct{{
		key:   2,
		value: "a",
	}}) {
		t.Error("GroupByConcurrentlyWithError returned unexpected value", result)
	}

	if !reflect.DeepEqual(result[4], []keyStruct{{
		key:   4,
		value: "b",
	}}) {
		t.Error("GroupByConcurrentlyWithError returned unexpected value", result)
	}

	if !reflect.DeepEqual(sets.SetFromSlice(result[6]), sets.SetFromSlice([]keyStruct{
		{
			key:   6,
			value: "c",
		},
		{
			key:   6,
			value: "d",
		},
	})) {
		t.Error("GroupByConcurrentlyWithError returned unexpected value", result)
	}

	values = []keyStruct{
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

	result, err = GroupByConcurrentlyWithError(values, keyStructToKeyWithError)
	if err == nil {
		t.Error("GroupByConcurrentlyWithError should have returned an error provided by its keyGenerator")
	}

	if !reflect.DeepEqual(result, map[int][]keyStruct{
		2: {{
			key:   2,
			value: "b",
		}},
	}) {
		t.Error("GroupByConcurrentlyWithError returned unexpected value", result)
	}
}
