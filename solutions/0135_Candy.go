package solutions

func Candy(ratings []int) int {
	res := 1
	slope := 0
	addend := 1
	culminating := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] == ratings[i-1] {
			if slope < 0 {
				res -= min(addend, culminating) - 1
			}
			addend = 1
		} else if ratings[i-1] < ratings[i] {
			if slope <= 0 {
				res -= min(addend, culminating) - 1
				addend = 1
			}
			addend++
		} else {
			if slope >= 0 {
				culminating = addend
				addend = 1
			}
			addend++
		}
		res += addend
		slope = ratings[i] - ratings[i-1]
	}
	if slope < 0 {
		res -= min(addend, culminating) - 1
		addend = 1
	}
	return res
}
