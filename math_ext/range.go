package math_ext

import "github.com/preston-wagner/unicycle/number"

func Range[N number.Number](start N, stop N, step N) []N {
	result := make([]N, 0, int((stop-start)/step))
	for i := start; i < stop; i += step {
		result = append(result, i)
	}
	return result
}
