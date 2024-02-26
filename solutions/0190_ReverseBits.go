package solutions

func ReverseBits(num uint32) uint32 {
	var res uint32 = 0
	high, low := 31, 0
	for high > low {
		res |= ((num >> high) & 1) << low
		res |= ((num >> low) & 1) << high
		high, low = high-1, low+1
	}
	return res
}
