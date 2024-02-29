package solutions

import "math"

func CountPrimes(n int) int {
	if n < 2 {
		return 0
	}
	cache := make([]int, n)
	cache[1] = 1
	upper := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < upper; i++ {
		if cache[i] != 0 {
			continue
		}
		for j := i * i; j < n; j += i {
			cache[j] = 1
		}
	}

	count := 0
	for i := 1; i < n; i++ {
		count += cache[i]
	}
	return n - count - 1
}
