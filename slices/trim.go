package slices

// Trim returns a copy of a slice with the underlying array shrunk to match the length of the used portion, or the original if the length and capacity are already the same
// Useful for preventing memory leaks; used internally by functions like Filter
// Performance: O(n)
func Trim[T any](input []T) []T {
	if len(input) == cap(input) {
		return input
	}
	trimmed := make([]T, len(input))
	copy(trimmed, input)
	return trimmed
}
