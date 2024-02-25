package solutions

func MajorityElement(nums []int) int {
	majority := 0
	voteNum := 0
	for _, num := range nums {
		if voteNum == 0 {
			majority = num
			voteNum++
		} else if num == majority {
			voteNum++
		} else {
			voteNum--
		}
	}
	return majority
}
