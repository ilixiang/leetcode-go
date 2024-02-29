package solutions

func IsIsomorphic(s string, t string) bool {
	m := len(s)
	mapping := map[byte]byte{}
	reverse := map[byte]byte{}

	for i := 0; i < m; i++ {
		c1, e1 := mapping[s[i]]
		c2, e2 := reverse[t[i]]
		if !(e1 || e2) {
			mapping[s[i]] = t[i]
			reverse[t[i]] = s[i]
		} else if e1 && e2 && c1 == t[i] && c2 == s[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
