package solutions

func strStr(haystack string, needle string) int {
	lps := make([]int, len(needle))
	for i := 2; i < len(lps); i++ {
		prev := lps[i-1]
		if needle[i-1] == needle[prev] {
			lps[i] = prev + 1
		} else {
			for j := prev; j > 0; j-- {
				k := 0
				for k < j && needle[k] == needle[i-j+k] {
					k++
				}
				if k == j {
					lps[i] = j
					break
				}
			}
		}
	}

	i, j := 0, 0
	for i <= len(haystack)-len(needle) {
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i
		} else if j != 0 {
			i = i + j - lps[j]
			j = lps[j]
		} else {
			i++
		}
	}
	return -1
}
