package unicycle

import (
	"errors"
	"testing"
)

type JsonPlaceholder struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TestFetchJson(t *testing.T) {
	placeholder, err := FetchJson[JsonPlaceholder]("https://jsonplaceholder.typicode.com/todos/1", FetchOptions{})
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
	_, err := FetchJson[JsonPlaceholder]("https://jsonplaceholder.typicode.com/", FetchOptions{})
	if err == nil {
		t.Error("Non-json response did not return error")
	}
	if fetchError := ErrorAs[FetchError](err); fetchError == nil {
		t.Error("FetchJson should have responded with an instance of FetchError (according to ErrorAs)")
	}
	var fetchError FetchError
	if !errors.As(err, &fetchError) {
		t.Error("FetchJson should have responded with an instance of FetchError (according to errors.As)")
	}
}

func TestFetchJsonWith404(t *testing.T) {
	_, err := FetchJson[JsonPlaceholder]("https://www.google.com/badUrl", FetchOptions{})
	if err == nil {
		t.Error("404 response did not return error")
	}
	if fetchError := ErrorAs[FetchError](err); fetchError == nil {
		t.Error("FetchJson should have responded with an instance of FetchError (according to ErrorAs)")
	}
	var fetchError FetchError
	if !errors.As(err, &fetchError) {
		t.Error("FetchJson should have responded with an instance of FetchError (according to errors.As)")
	}
}
