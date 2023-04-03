package unicycle

func sliceToChannel[INPUT_TYPE any](input []INPUT_TYPE) chan INPUT_TYPE {
	output := make(chan INPUT_TYPE, cap(input))
	for _, value := range input {
		output <- value
	}
	close(output)
	return output
}
