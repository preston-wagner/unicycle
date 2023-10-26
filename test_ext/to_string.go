package test_ext

import (
	"errors"
	"fmt"
)

func ToString(input int) string {
	return fmt.Sprintf("%d", input)
}

func ToStringErrIfNegative(input int) (string, error) {
	if input < 0 {
		return "", errors.New("toStringIfOddErrIfNegative(): negative number")
	}
	return fmt.Sprintf("%d", input), nil
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
