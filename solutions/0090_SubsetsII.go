package solutions

import (
	"math"
	"sort"
)

func SubsetsWithDup(nums []int) [][]int {
	size := len(nums)
	results := make([][]int, 1, int(math.Exp2(float64(size))))
	results[0] = make([]int, 0)

	sort.Ints(nums)
	buf := make([]int, 0, size)
	var combination func(int, int)
	combination = func(start int, pos int) {
		if pos == len(buf) {
			result := make([]int, len(buf))
			copy(result, buf)
			results = append(results, result)
			return
		}

		bound := len(nums) - len(buf) + pos + 1
		for i := start; i < bound; i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			buf[pos] = nums[i]
			combination(i+1, pos+1)
		}
	}

	for i := 1; i <= size-1; i++ {
		buf = append(buf, 0)
		combination(0, 0)
	}

	results = append(results, nums)
	return results
}
