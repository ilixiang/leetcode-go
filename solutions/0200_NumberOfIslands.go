package solutions

func NumIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(row int, col int) {
		if grid[row][col] != '1' {
			return
		}
		grid[row][col] = 'v'
		if row < m-1 {
			dfs(row+1, col)
		}
		if col < n-1 {
			dfs(row, col+1)
		}
		if row > 0 {
			dfs(row-1, col)
		}
		if col > 0 {
			dfs(row, col-1)
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == '1' {
				res++
				dfs(row, col)
			}
		}
	}
	return res
}
