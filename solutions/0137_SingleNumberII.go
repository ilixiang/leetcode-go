package solutions

func SingleNumberII(nums []int) int {
	var res int32 = 0
	for i := 0; i < 32; i++ {
		tmp := 0
		for _, num := range nums {
			tmp += (num >> i) & 1
		}
		if tmp%3 == 1 {
			res |= 1 << i
		}
	}
	return int(res)
}
