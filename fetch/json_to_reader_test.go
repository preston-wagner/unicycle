package fetch

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"

	"github.com/preston-wagner/unicycle/error_ext"
)

func TestJsonToReader(t *testing.T) {
	original := jsonPlaceholder{
		UserId:    943563,
		ID:        3425932,
		Title:     "king",
		Completed: true,
	}

	duplicated, err := ReadJson[jsonPlaceholder](JsonToReader(original))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(original, duplicated) {
		t.Errorf("JsonToReader/ReadJson error")
	}

	_, err = ReadJson[jsonPlaceholder](JsonToReader(make(chan int)))
	if err == nil {
		t.Error("attempting to marshal an invalid type should result in an error")
	}
	if _, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Error("attempting to marshal an invalid type should result in an UnsupportedTypeError")
	}

	_, err = ReadJson[jsonPlaceholder](JsonToReader(math.Inf(1)))
	if err == nil {
		t.Error("attempting to marshal an invalid value should result in an error")
	}
	if _, ok := err.(*json.UnsupportedValueError); !ok {
		t.Error("attempting to marshal an invalid value should result in an UnsupportedValueError, got:", err)
	}
}

func TestJsonFetchBody(t *testing.T) {
	_, err := Fetch("https://www.google.com/", FetchOptions{
		Body: JsonToReader(make(chan int)),
	})
	if err == nil {
		t.Error("calling Fetch with an unmarshallable body should result in an error")
	}
	if wrappedErr := error_ext.ErrorAs[json.UnsupportedTypeError](err); wrappedErr == nil {
		t.Error("calling Fetch with an unmarshallable body should result in an UnsupportedTypeError, got:", err)
	}
}
