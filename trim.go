package unicycle

// Trim shrinks a slice's underlying array to match the length of the used portion
func Trim[T any](input []T) []T {
	return append(make([]T, 0, len(input)), input...)
}
