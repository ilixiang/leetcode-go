package solutions

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}

	letterMap := map[byte][]byte{
		'1': nil,
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
		'0': nil,
	}

	rev := make([]string, 0, 4*len(digits))
	buf := make([]byte, len(digits))
	var combine func(int)
	combine = func(i int) {
		if i == len(digits) {
			rev = append(rev, string(buf))
			return
		}
		for _, letter := range letterMap[digits[i]] {
			buf[i] = letter
			combine(i + 1)
		}
	}
	combine(0)
	return rev
}
