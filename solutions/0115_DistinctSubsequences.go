package solutions

func NumDistinct(s string, t string) int {
	m, n := len(s), len(t)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	var dp func(int, int) int
	dp = func(i int, j int) int {
		if j == n {
			return 1
		}
		if i == m {
			return 0
		}

		cached := cache[i][j]
		if cached != -1 {
			return cached
		}

		cached = 0
		if m-i >= n-j {
			cached = dp(i+1, j)
			if s[i] == t[j] {
				cached += dp(i+1, j+1)
			}
		}
		cache[i][j] = cached
		return cached
	}
	return dp(0, 0)
}
