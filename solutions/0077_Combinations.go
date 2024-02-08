package solutions

func Combine(n int, k int) [][]int {
	size := 1
	for i := n; i > 1; i-- {
		size *= i
	}
	for i := k; i > 1; i-- {
		size /= i
	}
	for i := n - k; i > 1; i-- {
		size /= i
	}

	buf := make([]int, k)
	results := make([][]int, 0, size)
	var recursive func(int, int)
	recursive = func(start int, pos int) {
		if pos == k {
			result := make([]int, k)
			copy(result, buf)
			results = append(results, result)
			return
		}

		end := n + pos - k + 1
		for i := start; i <= end; i++ {
			buf[pos] = i
			recursive(i+1, pos+1)
		}
	}

	recursive(1, 0)
	return results
}
