package fetch

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/preston-wagner/unicycle/error_ext"
	"github.com/preston-wagner/unicycle/json_ext"
	"github.com/preston-wagner/unicycle/test_ext"
)

func TestFetchJson(t *testing.T) {
	placeholder, err := FetchJson[test_ext.JsonPlaceholder]("https://jsonplaceholder.typicode.com/todos/1", FetchOptions{})
	if err != nil {
		t.Error("Error fetching test json")
	}
	if placeholder.UserId != 1 {
		t.Error("placeholder.UserId was not expected value")
	}
	if placeholder.ID != 1 {
		t.Error("placeholder.ID was not expected value")
	}
	if placeholder.Title != "delectus aut autem" {
		t.Error("placeholder.Title was not expected value")
	}
	if placeholder.Completed != false {
		t.Error("placeholder.Completed was not expected value")
	}
}

func TestFetchJsonWithoutJson(t *testing.T) {
	_, err := FetchJson[test_ext.JsonPlaceholder]("https://jsonplaceholder.typicode.com/", FetchOptions{})
	if err == nil {
		t.Error("Non-json response did not return error")
	}
	if fetchError := error_ext.ErrorAs[FetchError](err); fetchError == nil {
		t.Error("FetchJson should have responded with an instance of FetchError (according to ErrorAs)")
	}
	var fetchError FetchError
	if !errors.As(err, &fetchError) {
		t.Error("FetchJson should have responded with an instance of FetchError (according to errors.As)")
	}
}

func TestFetchJsonWith404(t *testing.T) {
	_, err := FetchJson[test_ext.JsonPlaceholder]("https://www.google.com/badUrl", FetchOptions{})
	if err == nil {
		t.Error("404 response did not return error")
	}
	if !IsBadResponseWithCode(err, 404) {
		t.Error("404 response did not return error wrapping BadResponseError")
	}
	if fetchError := error_ext.ErrorAs[FetchError](err); fetchError == nil {
		t.Error("FetchJson should have responded with an instance of FetchError (according to ErrorAs)")
	}
	var fetchError FetchError
	if !errors.As(err, &fetchError) {
		t.Error("FetchJson should have responded with an instance of FetchError (according to errors.As)")
	}
}

func TestFetchJsonAlwaysWith400(t *testing.T) {
	type fbErrorResponse struct {
		Error struct {
			Code int
		}
	}
	response, err := FetchJson[fbErrorResponse]("https://graph.facebook.com/v9.0/me", FetchOptions{
		Query: map[string]string{
			"access_token": "fakeTokenForUnitTest",
		},
		AcceptBadResponse: true,
	})
	if err != nil {
		t.Error(err)
	}
	if response.Error.Code != 190 {
		t.Error("FetchJsonAlways https://graph.facebook.com/v8.0/me response.Error.Code != 190")
	}
}

func TestJsonFetchBody(t *testing.T) {
	_, err := Fetch("https://www.google.com/", FetchOptions{
		Body: json_ext.JsonToReader(make(chan int)),
	})
	if err == nil {
		t.Error("calling Fetch with an unmarshallable body should result in an error")
	}
	if wrappedErr := error_ext.ErrorAs[json.UnsupportedTypeError](err); wrappedErr == nil {
		t.Error("calling Fetch with an unmarshallable body should result in an UnsupportedTypeError, got:", err)
	}
}
