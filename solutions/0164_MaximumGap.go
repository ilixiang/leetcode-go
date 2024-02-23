package solutions

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	minNum, maxNum := 1000000001, -1
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}

	size := len(nums)
	diff := maxNum - minNum
	if diff == 0 {
		return 0
	}
	interval := diff / size
	if diff%size != 0 {
		interval++
	}

	buckets := make([][2]int, size+1)
	for i := range buckets {
		buckets[i] = [2]int{-1, -1}
	}

	for _, num := range nums {
		idx := (num - minNum) / interval
		bucket := buckets[idx]
		left, right := bucket[0], bucket[1]
		if left == -1 {
			buckets[idx] = [2]int{num, num}
		} else {
			buckets[idx] = [2]int{min(left, num), max(right, num)}
		}
	}

	res := 0
	prev := minNum
	for _, bucket := range buckets {
		if bucket[0] != -1 {
			res = max(res, bucket[0]-prev)
			prev = bucket[1]
		}
	}
	return res
}
