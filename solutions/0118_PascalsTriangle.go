package solutions

func Generate(numRows int) [][]int {
	results := make([][]int, 0, numRows)
	buf := make([]int, numRows)
	for row := 0; row < numRows; row++ {
		buf[row] = 1
		for col := row - 1; col > 0; col-- {
			buf[col] = buf[col] + buf[col-1]
		}
		result := make([]int, row+1)
		copy(result, buf[:row+1])
		results = append(results, result)
	}
	return results
}
