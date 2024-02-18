package solutions

func SingleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
