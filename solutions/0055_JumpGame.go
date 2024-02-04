package solutions

func CanJump(nums []int) bool {
	covered := nums[0]
	searched := 0
	for covered < len(nums)-1 && searched <= covered {
		curCovery := nums[searched] + searched
		if curCovery > covered {
			covered = curCovery
		}
		searched++
	}
	return covered >= len(nums)-1
}
