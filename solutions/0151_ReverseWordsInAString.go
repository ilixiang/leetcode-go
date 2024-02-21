package solutions

func ReverseWords(s string) string {
	m := len(s)
	bs := []byte(s)
	divider := 0
	for i := 0; i < m; i++ {
		if (divider != 0 && bs[divider-1] != ' ') || bs[i] != ' ' {
			bs[divider] = bs[i]
			divider++
		}
	}
	if bs[divider-1] == ' ' {
		divider--
	}

	i := 0
	for i < divider {
		for i < divider && bs[i] == ' ' {
			i++
		}
		left := i
		for i < divider && bs[i] != ' ' {
			i++
		}
		right := i - 1

		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}

	left, right := 0, divider-1
	for left < right {
		bs[left], bs[right] = bs[right], bs[left]
		left++
		right--
	}
	return string(bs[:divider])
}
