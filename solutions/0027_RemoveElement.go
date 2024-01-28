package solutions

func RemoveElement(nums []int, val int) int {
	idx := 0
	for _, num := range nums {
		if num != val {
			nums[idx] = num
			idx++
		}
	}
	return idx
}
