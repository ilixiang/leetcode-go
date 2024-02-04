package solutions

func Permute(nums []int) [][]int {
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

		for i := 0; i < len(nums); i++ {
			if accessed[i] == 0 {
				accessed[i] = 1
				buf[pos] = nums[i]
				recursive(pos + 1)
				accessed[i] = 0
			}
		}
	}

	recursive(0)
	return results
}
