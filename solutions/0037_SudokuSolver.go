package solutions

func solveSudoku(board [][]byte) {
	rows := [9]int{}
	cols := [9]int{}
	subs := [9]int{}

	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			if board[rowIdx][colIdx] != '.' {
				num := int(board[rowIdx][colIdx] - '0')
				mask := 1 << num
				subIdx := rowIdx/3*3 + colIdx/3
				rows[rowIdx] |= mask
				cols[colIdx] |= mask
				subs[subIdx] |= mask
			}
		}
	}

	var solve func(int) bool
	solve = func(idx int) bool {
		if idx == 81 {
			return true
		}
		rowIdx := idx / 9
		colIdx := idx % 9
		if board[rowIdx][colIdx] != '.' {
			return solve(idx + 1)
		} else {
			for num := 1; num < 10; num++ {
				mask := 1 << num
				subIdx := rowIdx/3*3 + colIdx/3
				if rows[rowIdx]&mask == 0 && cols[colIdx]&mask == 0 && subs[subIdx]&mask == 0 {
					board[rowIdx][colIdx] = '0' + byte(num)
					rows[rowIdx] |= mask
					cols[colIdx] |= mask
					subs[subIdx] |= mask
					if solve(idx + 1) {
						return true
					}
					board[rowIdx][colIdx] = '.'
					rows[rowIdx] &= ^mask
					cols[colIdx] &= ^mask
					subs[subIdx] &= ^mask
				}
			}
		}
		return false
	}

	solve(0)
}
