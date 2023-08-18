package unicycle

import (
	"reflect"
	"strconv"
	"testing"
)

func matches(key string, value int) bool {
	return strconv.Itoa(value) == key
}

func TestFilterMap(t *testing.T) {
	testData := map[string]int{
		"1": 1,
		"B": 2,
		"3": 3,
		"D": 4,
		"5": 5,
	}
	result := FilterMap(testData, matches)
	if !reflect.DeepEqual(result, map[string]int{
		"1": 1,
		"3": 3,
		"5": 5,
	}) {
		t.Errorf("FilterMap() returned unexpected %v", result)
	}

	if len(FilterMap(nil, matches)) != 0 {
		t.Error("FilterMap(nil) should return a map with length 0")
	}
}
