package solutions

func Jump(nums []int) int {
	start := 0
	covered := 0
	step := 0
	for covered < len(nums)-1 {
		tmp := covered
		for i := start; i <= tmp; i++ {
			reached := i + nums[i]
			if reached > covered {
				covered = reached
			}
		}
		start = tmp
		step++
	}
	return step
}
