package solutions

func HammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		res += int((num >> i) & 1)
	}
	return res
}
