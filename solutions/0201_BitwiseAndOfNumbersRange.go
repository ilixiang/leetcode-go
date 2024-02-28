package solutions

func RangeBitwiseAnd(left int, right int) int {
	res := 0
	tmp := 0
	diff := right - left
	for i := 0; i < 31; i++ {
		if (left>>i)&1 == 1 {
			tmp |= (1 << i) & left
			if diff < (1<<(i+1))-tmp {
				res |= 1 << i
			}
		}
	}
	res |= (1 << 31) & left
	return res
}
