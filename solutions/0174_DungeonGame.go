package solutions

import "math"

func CalculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	cache[m-1][n-1] = 1
	if dungeon[m-1][n-1] < 0 {
		cache[m-1][n-1] = 1 - dungeon[m-1][n-1]
	}

	var dp func(int, int) int
	dp = func(row int, col int) int {
		cached := cache[row][col]
		if cached != -1 {
			return cached
		}

		right := math.MaxInt32
		if col < n-1 {
			right = dp(row, col+1)
		}
		down := math.MaxInt32
		if row < m-1 {
			down = dp(row+1, col)
		}

		cached = max(1, -dungeon[row][col]+min(right, down))
		cache[row][col] = cached
		return cached
	}
	return dp(0, 0)
}
