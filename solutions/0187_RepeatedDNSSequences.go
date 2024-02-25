package solutions

func FindRepeatedDnaSequences(s string) []string {
	res := []string{}
	m := map[string]int{}
	for i := 0; i < len(s)-9; i++ {
		seq := s[i : i+10]
		count := m[seq]
		if count == 1 {
			res = append(res, seq)
		}
		m[seq] = count + 1
	}
	return res
}
