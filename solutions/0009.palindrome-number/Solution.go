package solutions

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]int, 0, 11)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	left, right := 0, len(digits)-1
	for left < right && digits[left] == digits[right] {
		left += 1
		right -= 1
	}
	return left >= right
}
