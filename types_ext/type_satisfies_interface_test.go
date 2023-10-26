package types_ext

import (
	"testing"
)

type getter interface {
	Get() string
}

type thingy struct{}

func (t thingy) Get() string {
	return "thingy"
}

type ptrThingy struct{}

func (t *ptrThingy) Get() string {
	return "ptrThingy"
}

type notThingy struct{}

func TestTypeSatisfiesInterface(t *testing.T) {
	if !TypeSatisfiesInterface[thingy, getter]() {
		t.Error("TypeSatisfiesInterface should return true when the provided type matches the interface")
	}
	if TypeSatisfiesInterface[ptrThingy, getter]() {
		t.Error("TypeSatisfiesInterface should return false when a pointer to the provided type matches the interface, but the type itself does not")
	}
	if TypeSatisfiesInterface[notThingy, getter]() {
		t.Error("TypeSatisfiesInterface should return false when the provided type does not match the interface")
	}
}
