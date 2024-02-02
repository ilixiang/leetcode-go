package solutions

func FirstMissingPositive(nums []int) int {
	i := 0
	length := len(nums)
	for i < length {
		num := nums[i]
		if num > 0 && num <= length && nums[num-1] != num {
			nums[num-1], nums[i] = nums[i], nums[num-1]
		} else {
			i++
		}
	}

	for i, num := range nums {
		if i != num-1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
