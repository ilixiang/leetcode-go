package solutions

func MinimumTotal(triangle [][]int) int {
	m := len(triangle)
	cache := make([][]int, m+1)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			cache[i][j] = -1
		}
	}
	cache[m] = make([]int, m+1)

	var dp func(int, int) int
	dp = func(row int, col int) int {
		cached := cache[row][col]
		if cached != -1 {
			return cached
		}

		cached = triangle[row][col] + min(dp(row+1, col), dp(row+1, col+1))
		cache[row][col] = cached
		return cached
	}
	return dp(0, 0)
}
