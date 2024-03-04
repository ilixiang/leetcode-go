package solutions

func RobII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := len(nums)
	cache := make([]int, m)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}
	var dp func(int, int) int
	dp = func(i int, end int) int {
		if i >= end {
			return 0
		}

		cached := cache[i]
		if cached != -1 {
			return cached
		}

		cached = max(nums[i]+dp(i+2, end), dp(i+1, end))
		cache[i] = cached
		return cached
	}

	res := dp(0, m-1)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}
	res = max(res, dp(1, m))
	return res
}
