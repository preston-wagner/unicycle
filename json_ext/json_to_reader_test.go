package json_ext

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/test_ext"
)

func TestJsonToReader(t *testing.T) {
	original := test_ext.JsonPlaceholder{
		UserId:    943563,
		ID:        3425932,
		Title:     "king",
		Completed: true,
	}

	duplicated, err := ReadJson[test_ext.JsonPlaceholder](JsonToReader(original))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(original, duplicated) {
		t.Errorf("JsonToReader/ReadJson error")
	}

	_, err = ReadJson[test_ext.JsonPlaceholder](JsonToReader(make(chan int)))
	if err == nil {
		t.Error("attempting to marshal an invalid type should result in an error")
	}
	if _, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Error("attempting to marshal an invalid type should result in an UnsupportedTypeError")
	}

	_, err = ReadJson[test_ext.JsonPlaceholder](JsonToReader(math.Inf(1)))
	if err == nil {
		t.Error("attempting to marshal an invalid value should result in an error")
	}
	if _, ok := err.(*json.UnsupportedValueError); !ok {
		t.Error("attempting to marshal an invalid value should result in an UnsupportedValueError, got:", err)
	}
}
