package solutions

func MinWindow(s string, t string) string {
	NOT_EXISTED_VAL := -100001

	charMap := make([]int, int('z'-'A')+1)
	for _, b := range t {
		charMap[int(b-'A')]++
	}

	total := 0
	for i := range charMap {
		if charMap[i] == 0 {
			charMap[i] = NOT_EXISTED_VAL
		} else {
			total++
		}
	}

	start, length := -1, len(s)+1
	left, right := 0, 0
	for right < len(s) || total == 0 {
		if total == 0 {
			if right-left < length {
				start = left
				length = right - left
			}
			idx := int(s[left] - 'A')
			if charMap[idx] != NOT_EXISTED_VAL {
				charMap[idx]++
				if charMap[idx] == 1 {
					total++
				}
			}
			left++
		} else {
			idx := int(s[right] - 'A')
			if charMap[idx] != NOT_EXISTED_VAL {
				charMap[idx]--
				if charMap[idx] == 0 {
					total--
				}
			}
			right++
		}
	}
	if start == -1 {
		return ""
	} else {
		return s[start : start+length]
	}
}
