package solutions

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		diff := target - num
		j, ok := m[diff]
		if ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
