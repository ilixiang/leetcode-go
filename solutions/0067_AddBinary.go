package solutions

func AddBinary(a string, b string) string {
	buf := make([]byte, 0, len(a)+len(b))
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 && j >= 0 {
		s := int(a[i]-'0'+b[j]-'0') + carry
		buf = append(buf, byte(s%2+'0'))
		carry = s / 2
		i--
		j--
	}
	for i >= 0 {
		s := int(a[i]-'0') + carry
		buf = append(buf, byte(s%2+'0'))
		carry = s / 2
		i--
	}
	for j >= 0 {
		s := int(b[j]-'0') + carry
		buf = append(buf, byte(s%2+'0'))
		carry = s / 2
		j--
	}
	if carry != 0 {
		buf = append(buf, byte(carry+'0'))
	}

	left, right := 0, len(buf)-1
	for left < right {
		buf[left], buf[right] = buf[right], buf[left]
		left++
		right--
	}
	return string(buf)
}
