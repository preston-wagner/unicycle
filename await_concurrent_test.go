package unicycle

import (
	"errors"
	"testing"
	"time"
)

func TestAwaitConcurrent(t *testing.T) {
	var a int
	var b int
	var c int

	start := time.Now()

	AwaitConcurrent(
		func() {
			time.Sleep(time.Second * 2)
			a = 1
		},
		func() {
			time.Sleep(time.Second * 1)
			b = 2
		},
		func() {
			time.Sleep(time.Second * 3)
			c = 3
		},
	)

	end := time.Now()

	if a != 1 {
		t.Error("concurrent function result not set!")
	}
	if b != 2 {
		t.Error("concurrent function result not set!")
	}
	if c != 3 {
		t.Error("concurrent function result not set!")
	}
	if end.Sub(start) > (time.Second * 4) {
		t.Error("concurrent")
	}
}

func TestAwaitConcurrentWithErrors(t *testing.T) {
	var a int
	var b int
	var c int

	start := time.Now()

	AwaitConcurrentWithErrors(
		func() error {
			time.Sleep(time.Second * 2)
			a = 1
			return nil
		},
		func() error {
			time.Sleep(time.Second * 1)
			b = 2
			return nil
		},
		func() error {
			time.Sleep(time.Second * 3)
			c = 3
			return nil
		},
	)

	end := time.Now()

	if a != 1 {
		t.Error("concurrent function result not set!")
	}
	if b != 2 {
		t.Error("concurrent function result not set!")
	}
	if c != 3 {
		t.Error("concurrent function result not set!")
	}
	if end.Sub(start) > (time.Second * 4) {
		t.Error("concurrent")
	}

	testErr := errors.New("test")
	err := AwaitConcurrentWithErrors(
		func() error {
			return nil
		},
		func() error {
			time.Sleep(time.Second * 1)
			return testErr
		},
		func() error {
			time.Sleep(time.Second * 2)
			return nil
		},
	)
	if err != testErr {
		t.Error("AwaitConcurrentWithErrors should return an error if one was returned from a wrapped function")
	}
}
