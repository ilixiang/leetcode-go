package solutions

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	cache := make([][]int, m*n)
	for row := 0; row < m; row++ {
		cache[row] = make([]int, n)
		for col := 0; col < n; col++ {
			if obstacleGrid[row][col] == 1 {
				cache[row][col] = 0
			} else {
				cache[row][col] = -1
			}
		}
	}
	cache[m-1][n-1] = 1

	var dp func(int, int) int
	dp = func(row int, col int) int {
		cached := cache[row][col]
		if cached != -1 {
			return cached
		}

		val := 0
		if row < m-1 {
			val += dp(row+1, col)
		}
		if col < n-1 {
			val += dp(row, col+1)
		}
		cache[row][col] = val
		return val
	}
	return dp(0, 0)
}
