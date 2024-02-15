package solutions

func ParlindromePartition(s string) [][]string {
	m := len(s)
	parlindromeCache := make([][]int, m)
	for i := 0; i < m; i++ {
		parlindromeCache[i] = make([]int, m+1)
	}

	var isParlindrome func(int, int) bool
	isParlindrome = func(left int, right int) bool {
		cached := parlindromeCache[left][right]
		if cached != 0 {
			return cached == 1
		}

		cached = -1
		if left >= right || (s[left] == s[right] && isParlindrome(left+1, right-1)) {
			cached = 1
		}
		parlindromeCache[left][right] = cached
		return cached == 1
	}

	candidates := make([][]int, m)
	for left := 0; left < m; left++ {
		for right := left; right < m; right++ {
			if isParlindrome(left, right) {
				candidates[left] = append(candidates[left], right)
			}
		}
	}

	res := [][]string{}
	buf := []string{}
	var backtrack func(int)
	backtrack = func(i int) {
		if i == m {
			tmp := make([]string, len(buf))
			copy(tmp, buf)
			res = append(res, tmp)
			return
		}

		for _, right := range candidates[i] {
			buf = append(buf, s[i:right+1])
			backtrack(right + 1)
			buf = buf[:len(buf)-1]
		}
	}
	return res
}
