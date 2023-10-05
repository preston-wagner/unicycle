package math

// returns the average of the values of the provided array (or 0 if the array is empty)
func Average[N Number](input []N) N {
	count := len(input)
	if count > 0 {
		return Sum(input) / N(count)
	}
	return 0
}

// returns the average of the values of the provided array (or 0 if the array is empty) as a float64
func Average64[N Number](input []N) float64 {
	count := len(input)
	if count > 0 {
		return float64(Sum(input)) / float64(count)
	}
	return 0
}
