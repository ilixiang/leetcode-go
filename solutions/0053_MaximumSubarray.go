package solutions

func MaxSubArray(nums []int) int {
	minSum := 0
	sum := 0
	rev := nums[0]
	for _, num := range nums {
		sum += num
		if sum-minSum > rev {
			rev = sum - minSum
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return rev
}
