package solutions

func longestSubstringWithoutRepeatCharacters(s string) int {
	m := map[rune]int{}
	start := -1
	rev := 0
	for i, c := range s {
		exists, ok := m[c]
		if ok && exists > start {
			rev = max(rev, i-start-1)
			start = exists
		}
		m[c] = i
	}
	rev = max(rev, len(s)-start-1)
	return rev
}
