package unicycle

// SliceRefs accepts a slice of data and returns a new slice of references to the original slice values
func SliceRefs[INPUT_TYPE any](input []INPUT_TYPE) []*INPUT_TYPE {
	output := make([]*INPUT_TYPE, len(input))
	for index := range input {
		output[index] = &input[index]
	}
	return output
}
