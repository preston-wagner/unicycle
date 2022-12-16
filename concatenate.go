package unicycle

// Concatenate takes any number of slices and copies their values into a new slice
func Concatenate[INPUT_TYPE any](input ...[]INPUT_TYPE) []INPUT_TYPE {
	resultLength := Reduce(input, func(acc int, arr []INPUT_TYPE) int {
		return acc + len(arr)
	}, 0)
	result := make([]INPUT_TYPE, resultLength)
	start := 0
	for _, value := range input {
		for index2, value2 := range value {
			result[start+index2] = value2
		}
		start += len(value)
	}
	return result
}
