package unicycle

import (
	"errors"
	"testing"
)

type wrappingError struct {
	err error
}

func (e wrappingError) Error() string {
	return e.err.Error()
}

func (e wrappingError) Unwrap() error {
	return e.err
}

func TestErrorAs(t *testing.T) {
	var err error

	if fetchErr := ErrorAs[FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a nil error")
	}
	err = errors.New("irrelevant")
	if fetchErr := ErrorAs[FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given an instantiated error of a non-wrapping type")
	}
	err = FetchError{Err: errors.New("irrelevant")}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given an instance of it")
	}
	err = &FetchError{Err: errors.New("irrelevant")}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to an instance of it")
	}
	err = wrappingError{err: FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a type that wraps it")
	}
	err = &wrappingError{err: FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to a type that wraps it")
	}
	err = wrappingError{err: &FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a type that wraps a pointer to it")
	}
	err = &wrappingError{err: &FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := ErrorAs[FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to a type that wraps a pointer to it")
	}
	err = wrappingError{err: errors.New("irrelevant")}
	if fetchErr := ErrorAs[FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a pointer to a type that does not wrap FetchError")
	}
	err = &wrappingError{err: errors.New("irrelevant")}
	if fetchErr := ErrorAs[FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a pointer to a type that does not wrap FetchError")
	}
}
