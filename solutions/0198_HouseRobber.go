package solutions

func Rob(nums []int) int {
	m := len(nums)
	cache := make([]int, m)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}

	var dp func(int) int
	dp = func(i int) int {
		if i >= m {
			return 0
		}

		cached := cache[i]
		if cached != -1 {
			return cached
		}
		cached = max(nums[i]+dp(i+2), dp(i+1))
		cache[i] = cached
		return cached
	}
	return dp(0)
}
