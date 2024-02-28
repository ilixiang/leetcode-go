package solutions

func IsHappy(n int) bool {
	presenceMap := map[int]struct{}{}
	existed := false
	for !existed && n != 1 {
		presenceMap[n] = struct{}{}
		next := 0
		for n != 0 {
			remainder := n % 10
			next += remainder * remainder
			n /= 10
		}
		n = next
		_, existed = presenceMap[n]
	}
	return n == 1
}
