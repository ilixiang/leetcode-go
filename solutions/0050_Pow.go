package solutions

func MyPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 1 / MyPow(x, -n)
	} else {
		prev := MyPow(x, n>>1)
		if n&1 == 1 {
			return prev * prev * x
		} else {
			return prev * prev
		}
	}
}
