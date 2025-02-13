package error_ext_test

import (
	"errors"
	"testing"

	"github.com/preston-wagner/unicycle/error_ext"
)

func TestWrapError(t *testing.T) {
	var err error

	err = error_ext.WrapError("test: ", err)

	if err != nil {
		t.Error("WrapError should return nil when passed nil")
	}

	err = errors.New("An actual error this time")

	err = error_ext.WrapError("test: ", err)

	if err == nil {
		t.Error("WrapError should return an error when passed one")
	}

	if err.Error() != "test: An actual error this time" {
		t.Error("WrapError incorrectly joined the prefix and error message")
		t.Log(err.Error())
	}
}
