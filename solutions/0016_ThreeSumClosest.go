package solutions

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	closest := nums[0] + nums[1] + nums[2]
	closestDiff := target - closest
	if closestDiff < 0 {
		closestDiff = -1 * closestDiff
	}

	sort.Ints(nums)
	for idx := 0; idx < len(nums)-2; idx++ {
		if idx != 0 && nums[idx] == nums[idx-1] {
			continue
		}

		left, right := idx+1, len(nums)-1
		for left < right {
			sum := nums[idx] + nums[left] + nums[right]
			diff := target - sum
			if diff == 0 {
				return sum
			} else if diff < 0 {
				diff = -1 * diff
				right--
			} else {
				left++
			}

			if diff < closestDiff {
				closest = sum
				closestDiff = diff
			}
		}
	}
	return closest
}
