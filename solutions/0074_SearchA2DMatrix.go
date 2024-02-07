package solutions

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) >> 1
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
