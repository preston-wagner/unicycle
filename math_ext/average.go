package math_ext

import "github.com/nuvi/unicycle/number"

// returns the average of the provided values (or 0 if there are none)
func Average[N number.Number](input ...N) N {
	count := len(input)
	if count > 0 {
		return Sum(input...) / N(count)
	}
	return 0
}

// returns the average of the provided values (or 0 if there are none) as a float64 for maximum precision
func Average64[N number.Number](input ...N) float64 {
	count := len(input)
	if count > 0 {
		return float64(Sum(input...)) / float64(count)
	}
	return 0
}
