package solutions

func GetRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for row := 0; row <= rowIndex; row++ {
		result[row] = 1
		for col := row - 1; col > 0; col-- {
			result[col] = result[col] + result[col-1]
		}
	}
	return result
}
