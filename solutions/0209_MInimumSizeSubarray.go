package solutions

import "math"

func MinSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	left, right := 0, 0
	s := 0
	for right < len(nums) {
		if s >= target {
			res = min(res, right-left)
			s -= nums[left]
			left++
		} else {
			s += nums[right]
			right++
		}
	}

	for s >= target {
		res = min(res, right-left)
		left++
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}
