package solutions

func removeDuplicates(nums []int) int {
	idx := 0
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}
