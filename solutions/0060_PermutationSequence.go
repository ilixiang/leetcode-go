package solutions

func GetPermutation(n int, k int) string {
	buf := make([]byte, n)
	order := make([]int, n)
	posibility := 1
	for i := 0; i < n; i++ {
		posibility *= i + 1
		order[i] = i + 1
	}

	k -= 1
	for i := 0; i < n; i++ {
		posibility /= n - i
		idx := k / posibility
		buf[i] = byte(order[idx] + '0')
		order = append(order[:idx], order[idx+1:]...)
		k %= posibility
	}

	return string(buf)
}
