package solutions

import "strconv"

func CompareVersion(v1 string, v2 string) int {
	i1, i2 := -1, -1
	p1, p2 := 0, 0
	for p1 == p2 && (i1 < len(v1) || i2 < len(v2)) {
		i1, i2 = i1+1, i2+1
		p1, p2 = 0, 0
		s1, s2 := i1, i2
		for i1 < len(v1) && '0' <= v1[i1] && v1[i1] <= '9' {
			i1++
		}
		if i1 != s1 {
			p1, _ = strconv.Atoi(v1[s1:i1])
		}

		for i2 < len(v2) && '0' <= v2[i2] && v2[i2] <= '9' {
			i2++
		}
		if i2 != s2 {
			p2, _ = strconv.Atoi(v2[s2:i2])
		}
	}

	if p1 == p2 {
		return 0
	} else if p1 > p2 {
		return 1
	} else {
		return -1
	}
}
