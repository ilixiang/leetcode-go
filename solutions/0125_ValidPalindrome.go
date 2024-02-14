package solutions

func isAlphanumeric(b byte) bool {
	return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') || ('0' <= b && b <= '9')
}

func equalsIgnoreCase(a byte, b byte) bool {
	if 'A' <= a && a <= 'Z' {
		a = a - 'A' + 'a'
	}
	if 'A' <= b && b <= 'Z' {
		b = b - 'A' + 'a'
	}
	return a == b
}

func IsPalindromeIgnoreNonAlphabets(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
		} else if !isAlphanumeric(s[right]) {
			right--
		} else if !equalsIgnoreCase(s[left], s[right]) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
