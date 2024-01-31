package solutions

func NextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	idx := len(nums) - 1
	for idx > 0 && nums[idx] <= nums[idx-1] {
		idx--
	}

	left, right := idx, len(nums)-1
	if idx != 0 {
		targetNum := nums[idx-1]
		targetIdx := left
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] > targetNum {
				targetIdx = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		nums[targetIdx], nums[idx-1] = nums[idx-1], nums[targetIdx]
	}

	left, right = idx, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
