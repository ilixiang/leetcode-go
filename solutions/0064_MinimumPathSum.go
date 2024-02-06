package solutions

func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([][]int, m)
	for row := 0; row < m; row++ {
		cache[row] = make([]int, n)
		for col := 0; col < n; col++ {
			cache[row][col] = -1
		}
	}
	cache[m-1][n-1] = grid[m-1][n-1]

	var dp func(int, int) int
	dp = func(row int, col int) int {
		cached := cache[row][col]
		if cached != -1 {
			return cached
		}

		next := 40000
		if row < m-1 {
			val := dp(row+1, col)
			if next > val {
				next = val
			}
		}
		if col < n-1 {
			val := dp(row, col+1)
			if next > val {
				next = val
			}
		}
		cache[row][col] = next + grid[row][col]
		return cache[row][col]
	}
	return dp(0, 0)
}
