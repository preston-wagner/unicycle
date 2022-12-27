package unicycle

// Trim returns a copy of a slice with the underlying array shrunk to match the length of the used portion
func Trim[T any](input []T) []T {
	trimmed := make([]T, 0, len(input))
	copy(trimmed, input)
	return trimmed
}
