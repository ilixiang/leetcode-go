package solutions

func CanCompleteCircuit(gas []int, cost []int) int {
	m := len(gas)
	idx := 0
	minimumRemain := 1000
	remain := 0
	for i := 0; i < m; i++ {
		remain += gas[i] - cost[i]
		if minimumRemain >= remain {
			idx = i
			minimumRemain = remain
		}
	}
	if remain < 0 {
		return -1
	}
	return (idx + 1) % m
}
