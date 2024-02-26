package solutions

func RotateArray(nums []int, k int) {
	m := len(nums)
	k = k % m
	left, right := 0, m-k-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}

	left, right = m-k, m-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}

	left, right = 0, m-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}
}
