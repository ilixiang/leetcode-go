package solutions

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func ContainsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	window := redblacktree.NewWithIntComparator()
	for i := 0; i < len(nums); i++ {
		if i > indexDiff {
			window.Remove(nums[i-indexDiff-1])
		}

		num := nums[i]
		floor, existed := window.Floor(num)
		if existed && num-floor.Key.(int) <= valueDiff {
			return true
		}

		ceiling, existed := window.Ceiling(num)
		if existed && ceiling.Key.(int)-num <= valueDiff {
			return true
		}
		window.Put(nums[i], struct{}{})
	}
	return false
}
