package solutions

func MaxProduct(nums []int) int {
	res := -int(uint(1<<31)-1) - 1
	minNegative, maxPositive := 0, 0
	for _, num := range nums {
		if num == 0 {
			minNegative, maxPositive = 0, 0
		} else {
			n1, n2 := num*maxPositive, num*minNegative
			minNegative = min(num, n1, n2)
			maxPositive = max(num, n1, n2)
		}
		res = max(res, minNegative)
		if maxPositive != 0 {
			res = max(res, maxPositive)
		}
	}
	return res
}
