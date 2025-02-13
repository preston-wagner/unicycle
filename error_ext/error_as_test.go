package error_ext_test

import (
	"errors"
	"testing"

	"github.com/nuvi/unicycle/error_ext"
	"github.com/nuvi/unicycle/fetch"
)

type WrappingError struct {
	err error
}

func (e WrappingError) Error() string {
	return e.err.Error()
}

func (e WrappingError) Unwrap() error {
	return e.err
}

type PtrError struct {
	err error
}

func (e *PtrError) Error() string {
	return e.err.Error()
}

func (e *PtrError) Unwrap() error {
	return e.err
}

func TestErrorAs(t *testing.T) {
	var err error

	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a nil error")
	}
	err = errors.New("irrelevant")
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given an instantiated error of a non-wrapping type")
	}
	err = fetch.FetchError{Err: errors.New("irrelevant")}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given an instance of it")
	}
	err = &fetch.FetchError{Err: errors.New("irrelevant")}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to an instance of it")
	}

	err = WrappingError{err: fetch.FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a type that wraps it")
	}
	err = &WrappingError{err: fetch.FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to a type that wraps it")
	}
	err = WrappingError{err: &fetch.FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a type that wraps a pointer to it")
	}
	err = &WrappingError{err: &fetch.FetchError{Err: errors.New("irrelevant")}}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr == nil {
		t.Error("ErrorAs should have returned an instance of FetchError when given a pointer to a type that wraps a pointer to it")
	}
	err = WrappingError{err: errors.New("irrelevant")}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a pointer to a type that does not wrap FetchError")
	}
	err = &WrappingError{err: errors.New("irrelevant")}
	if fetchErr := error_ext.ErrorAs[fetch.FetchError](err); fetchErr != nil {
		t.Error("ErrorAs should have returned nil when given a pointer to a type that does not wrap FetchError")
	}

	err = &PtrError{err: errors.New("irrelevant")}
	if ptrErr := error_ext.ErrorAs[PtrError](err); ptrErr == nil {
		t.Error("ErrorAs should not have returned nil when given a pointer to the expected type")
	}
	err = WrappingError{err: &PtrError{err: errors.New("irrelevant")}}
	if ptrErr := error_ext.ErrorAs[PtrError](err); ptrErr == nil {
		t.Error("ErrorAs should have returned an instance of PtrError when given a type that wraps a pointer to it")
	}
	err = WrappingError{err: &PtrError{err: errors.New("irrelevant")}}
	if ptrErr := error_ext.ErrorAs[PtrError](err); ptrErr == nil {
		t.Error("ErrorAs should have returned an instance of ptrError when given a type that wraps a pointer to it")
	}
	err = &WrappingError{err: &PtrError{err: errors.New("irrelevant")}}
	if ptrErr := error_ext.ErrorAs[PtrError](err); ptrErr == nil {
		t.Error("ErrorAs should have returned an instance of ptrError when given a pointer to a type that wraps a pointer to it")
	}
}
