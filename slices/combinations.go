package slices

// Combinations accepts a slice of any data type, and returns a slice of slices representing all possible combinations of that type (preserving order of individual elements)
func Combinations[INPUT_TYPE any](input []INPUT_TYPE) [][]INPUT_TYPE {
	if len(input) == 0 {
		return [][]INPUT_TYPE{{}}
	} else {
		children := Combinations(input[1:])
		withCurrent := Mapping(children, func(child []INPUT_TYPE) []INPUT_TYPE {
			return Concatenate([]INPUT_TYPE{input[0]}, child)
		})
		return Concatenate(withCurrent, children)
	}
}
