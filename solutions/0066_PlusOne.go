package solutions

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i] + carry
		digits[i] = s % 10
		carry = s / 10
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
