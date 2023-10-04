package unicycle

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func toString(input int) string {
	return fmt.Sprintf("%d", input)
}

func toStringErrIfNegative(input int) (string, error) {
	if input < 0 {
		return "", errors.New("toStringIfOddErrIfNegative(): negative number")
	}
	return fmt.Sprintf("%d", input), nil
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

func TestMappingWithError(t *testing.T) {
	result, err := MappingWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, toStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingWithError() returned unexpected %s", result)
	}

	result, err = MappingWithError(nil, toStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingWithError(nil) should return a slice with length 0")
	}

	_, err = MappingWithError([]int{1, 2, 3, -1, 7, 8}, toStringErrIfNegative)
	if err == nil {
		t.Error("MappingWithError should return error if any mapping functions do")
	}
}

func TestMappingMultithread(t *testing.T) {
	result := MappingMultithread([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, toString)
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingMultithread() returned unexpected %s", result)
	}

	if len(MappingMultithread(nil, toString)) != 0 {
		t.Error("MappingMultithread(nil) should return a slice with length 0")
	}
}

func TestMappingMultithreadWithError(t *testing.T) {
	result, err := MappingMultithreadWithError([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, toStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}) {
		t.Errorf("MappingMultithreadWithError() returned unexpected %s", result)
	}

	result, err = MappingMultithreadWithError(nil, toStringErrIfNegative)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("MappingMultithreadWithError(nil) should return a slice with length 0")
	}

	_, err = MappingMultithreadWithError([]int{1, 2, 3, -1, 7, 8}, toStringErrIfNegative)
	if err == nil {
		t.Error("MappingMultithreadWithError should return error if any mapping functions do")
	}
}
