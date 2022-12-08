package unicycle

func Mapping[INPUT_TYPE any, OUTPUT_TYPE any](input []INPUT_TYPE, mapper func(INPUT_TYPE) OUTPUT_TYPE) []OUTPUT_TYPE {
	output := make([]OUTPUT_TYPE, len(input))
	for index, value := range input {
		output[index] = mapper(value)
	}
	return output
}
