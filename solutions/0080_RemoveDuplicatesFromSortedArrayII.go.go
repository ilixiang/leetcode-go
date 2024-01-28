package solutions

func RemoveDuplicatesII(nums []int) int {
	idx := 0
	count := 0
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			count += 1
		} else {
			count = 0
		}
		if count < 2 {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
