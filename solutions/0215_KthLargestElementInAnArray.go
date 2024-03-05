package solutions

import "math/rand"

func FindKthLargest(nums []int, k int) int {
	target := len(nums) - k
	partition := func(left int, right int) int {
		idx := rand.Intn(right-left+1) + left
		nums[left], nums[idx] = nums[idx], nums[left]

		start, end := left, left
		pivot := nums[left]
		for i := left + 1; i <= right; i++ {
			if nums[i] < pivot {
				nums[i], nums[start] = nums[start], nums[i]
				start, end = start+1, end+1
				nums[i], nums[end] = nums[end], nums[i]
			} else if pivot == nums[i] {
				end++
				nums[i], nums[end] = nums[end], nums[i]
			}
		}
		if start <= target && target <= end {
			return target
		} else if target < start {
			return start
		} else {
			return end
		}
	}

	left, right := 0, len(nums)-1
	idx := partition(left, right)
	for idx != target {
		if idx < target {
			left = idx + 1
		} else {
			right = idx - 1
		}
		idx = partition(left, right)
	}
	return nums[idx]
}
