package math_ext

// returns the largest of the provided values (or 0 if there are none)
func Max[N Number](input ...N) N {
	if len(input) == 0 {
		return 0
	} else {
		largest := input[0]
		for _, value := range input {
			if value > largest {
				largest = value
			}
		}
		return largest
	}
}
