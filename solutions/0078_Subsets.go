package solutions

import "math"

func Subsets(nums []int) [][]int {
	m := len(nums)
	size := int(math.Exp2(float64(m)))
	results := make([][]int, 0, size)

	buf := make([]int, m)
	var recursive func(int, int)
	recursive = func(start int, pos int) {
		if pos == len(buf) {
			result := make([]int, len(buf))
			copy(result, buf)
			results = append(results, result)
			return
		}

		end := m + pos - len(buf)
		for i := start; i <= end; i++ {
			buf[pos] = nums[i]
			recursive(i+1, pos+1)
		}
	}

	for i := m; i > 0; i-- {
		recursive(0, 0)
		buf = buf[:len(buf)-1]
	}
	results = append(results, []int{})
	return results
}
