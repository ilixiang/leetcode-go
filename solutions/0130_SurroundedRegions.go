package solutions

func Solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(row int, col int) {
		if board[row][col] == 'O' {
			board[row][col] = 'V'
			if row > 0 {
				dfs(row-1, col)
			}
			if col > 0 {
				dfs(row, col-1)
			}
			if row < m-1 {
				dfs(row+1, col)
			}
			if col < n-1 {
				dfs(row, col+1)
			}
		}
	}

	for row := 0; row < m; row++ {
		dfs(row, 0)
		dfs(row, n-1)
	}
	for col := 0; col < n; col++ {
		dfs(0, col)
		dfs(m-1, col)
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == 'V' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}
