package solutions

func SortColors(nums []int) {
	colors := [3]int{0, 0, 0}
	for _, num := range nums {
		colors[num] += 1
	}

	idx := 0
	for i := 0; i < 3; i++ {
		for count := colors[i]; count > 0; count-- {
			nums[idx] = i
			idx++
		}
	}
}
