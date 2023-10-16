package test_ext

import (
	"errors"
	"fmt"
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

func ToStringIfOdd(input int) (string, bool) {
	if Odd(input) {
		return fmt.Sprintf("%d", input), true
	}
	return "", false
}

func ToStringIfOddErrIfNegative(input int) (string, bool, error) {
	if input < 0 {
		return "", false, errors.New("toStringIfOddErrIfNegative(): negative number")
	}
	if Odd(input) {
		return fmt.Sprintf("%d", input), true, nil
	}
	return "", false, nil
}
