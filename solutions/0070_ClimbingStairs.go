package solutions

func ClimbStairs(n int) int {
	cache := make([]int, n+1)
	for i := n; i > 0; i-- {
		cache[i] = -1
	}
	cache[0] = 1
	cache[1] = 1

	var dp func(int) int
	dp = func(m int) int {
		cached := cache[m]
		if cached != -1 {
			return cached
		}

		cached = dp(m-1) + dp(m-2)
		cache[m] = cached
		return cached
	}
	return dp(n)
}
