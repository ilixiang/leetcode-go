package solutions

import "sort"

func PermuteUniq(nums []int) [][]int {
	sort.Ints(nums)

	capacity := 1
	for i := 2; i <= len(nums); i++ {
		capacity *= i
	}
	results := make([][]int, 0, capacity)
	accessed := make([]int, len(nums))

	buf := make([]int, len(nums))
	var recursive func(int)
	recursive = func(pos int) {
		if pos == len(nums) {
			result := make([]int, len(nums))
			copy(result, buf)
			results = append(results, result)
			return
		}

		prev := -11
		for i := 0; i < len(nums); i++ {
			if nums[i] != prev && accessed[i] == 0 {
				accessed[i] = 1
				buf[pos] = nums[i]
				recursive(pos + 1)
				accessed[i] = 0
				prev = nums[i]
			}
		}
	}

	recursive(0)
	return results
}
