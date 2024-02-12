package solutions

import "math"

func GrayCode(n int) []int {
	size := int(math.Exp2(float64(n)))
	result := make([]int, size)

	shift := 0
	for (1 << shift) < size {
		bound := 1 << shift
		for i := 0; i < bound; i++ {
			result[bound+i] = result[bound-i-1] | bound
		}
		shift++
	}
	return result
}
