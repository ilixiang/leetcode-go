package solutions

func ShortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	bs := make([]byte, 0, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		bs = append(bs, '*')
		bs = append(bs, s[i])
	}
	bs = append(bs, '*')

	m := len(bs)
	center, covered := 0, 0
	radiuses := make([]int, len(bs))
	for i := 0; i < m; i++ {
		if i >= covered {
			left, right := i-1, i+1
			for left >= 0 && right < m && bs[left] == bs[right] {
				left, right = left-1, right+1
			}
			radiuses[i] = right - 1 - i
			center, covered = i, right-1
		} else {
			symmetry := 2*center - i
			radius := radiuses[symmetry]
			if i+radius != covered {
				radiuses[i] = min(radius, covered-i)
			} else {
				left, right := i-radius-1, i+radius+1
				for left >= 0 && right < m && bs[left] == bs[right] {
					left, right = left-1, right+1
				}
				radiuses[i] = right - 1 - i
				center, covered = i, right-1
			}
		}
	}

	palindromeLen := 1
	for i := 0; i < m; i++ {
		radius := radiuses[i]
		if i-radius == 0 {
			palindromeLen = max(palindromeLen, radius)
		}
	}

	buff := make([]byte, 0, len(s)-palindromeLen)
	for i := len(s) - 1; i >= palindromeLen; i-- {
		buff = append(buff, s[i])
	}
	return string(buff) + s
}
