package unicycle

import (
	"errors"
	"reflect"
)

// ErrorAs simplifies checking the types of errors, specifically in if statements:
//
//	if custErr := ErrorAs[CustomErrorType](err); custErr != nil {
//		 // use custErr
//	}
//
// it also handles edge cases missed by errors.As, such as pointers to the checked or intermediate types, or types that don't satisfy the error interface while pointers to the type do
//
// (non-nil pointers to structs that implement an interface also implement that interface as far as go is concerned in all other instances)
func ErrorAs[ERROR_TYPE any](err error) *ERROR_TYPE {
	if err != nil {
		if TypeSatisfiesInterface[ERROR_TYPE, error]() {
			var anyInstance ERROR_TYPE
			if errors.As(err, &anyInstance) {
				return &anyInstance
			}
		}
		if TypeSatisfiesInterface[*ERROR_TYPE, error]() {
			ptr, ok := any(err).(*ERROR_TYPE)
			if ok {
				return ptr
			}
		}
		if dereferencedErr := unsafeErrorAs[ERROR_TYPE](err); dereferencedErr != nil {
			return dereferencedErr
		}
		if wrapped := errors.Unwrap(err); wrapped != nil {
			return ErrorAs[ERROR_TYPE](wrapped)
		}
	}
	return nil
}

func unsafeErrorAs[ERROR_TYPE any](err error) (returnedErr *ERROR_TYPE) {
	defer func() {
		// you really never know when reflect is going to panic on you
		if r := recover(); r != nil {
			returnedErr = nil
		}
	}()
	// if err is a pointer, dereference it and try again
	if reflectValue := reflect.ValueOf(err); reflectValue.Kind() == reflect.Ptr {
		elem := reflectValue.Elem()
		intr := elem.Interface()
		dereferencedErr, ok := intr.(error)
		if ok {
			return ErrorAs[ERROR_TYPE](dereferencedErr)
		}
	}
	return nil
}
