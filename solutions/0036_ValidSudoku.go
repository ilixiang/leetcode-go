package solutions

func IsValidSudoku(board [][]byte) bool {
	rows := [9]int{}
	cols := [9]int{}
	subs := [9]int{}
	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			if board[rowIdx][colIdx] == '.' {
				continue
			}
			num := int(board[rowIdx][colIdx] - '0')
			subIdx := rowIdx/3*3 + colIdx/3
			mask := 1 << num
			if rows[rowIdx]&mask != 0 || cols[colIdx]&mask != 0 || subs[subIdx]&mask != 0 {
				return false
			}
			rows[rowIdx] |= mask
			cols[colIdx] |= mask
			subs[subIdx] |= mask
		}
	}
	return true
}
