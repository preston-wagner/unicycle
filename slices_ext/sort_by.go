package slices_ext

import (
	"golang.org/x/exp/constraints"
	"slices"
)

// SortBy accepts a slice of data and key generator function, and returns a slice of that data sorted by that key
func SortBy[KEY_TYPE constraints.Ordered, VALUE_TYPE any](input []VALUE_TYPE, keyGenerator func(VALUE_TYPE) KEY_TYPE) []VALUE_TYPE {
	sorted := make([]VALUE_TYPE, len(input))
	copy(sorted, input)
	slices.SortStableFunc(sorted, func(a VALUE_TYPE, b VALUE_TYPE) int {
		key_a := keyGenerator(a)
		key_b := keyGenerator(b)
		if key_a > key_b {
			return 1
		} else if key_a < key_b {
			return -1
		} else {
			return 0
		}
	})
	return sorted
}
