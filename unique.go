package unicycle

// Unique returns a copy of a slice with the duplicate values omitted, preserving order based on the first instance of each element
func Unique[T comparable](input []T) []T {
	set := Set[T]{}
	output := make([]T, 0, len(input))
	for _, value := range input {
		if !set.Has(value) {
			output = append(output, value)
			set.Add(value)
		}
	}
	return Trim(output)
}
