package math

// returns the sum of the values of the provided array (or 0 if the array is empty)
func Sum[N Number](input []N) N {
	var total N
	for _, value := range input {
		total += value
	}
	return total
}
