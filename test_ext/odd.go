package test_ext

import (
	"errors"
)

func Odd(input int) bool {
	return input%2 == 1
}

func OddErrIfNegative(input int) (bool, error) {
	if input < 0 {
		return false, errors.New("OddErrIfNegative(): negative number")
	}
	return input%2 == 1, nil
}
