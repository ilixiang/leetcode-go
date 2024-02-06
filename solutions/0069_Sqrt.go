package solutions

func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	if x == 2 || x == 3 {
		return 1
	}

	result := 0
	left, right := 2, x/2
	for left <= right {
		mid := (left + right) >> 1
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
