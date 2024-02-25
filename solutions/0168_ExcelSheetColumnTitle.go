package solutions

func ConvertToTitle(columnNumber int) string {
	buff := []byte{}
	for columnNumber != 0 {
		columnNumber--
		letter := byte((columnNumber % 26) + 'A')
		buff = append(buff, letter)
		columnNumber /= 26
	}

	left, right := 0, len(buff)-1
	for left < right {
		buff[left], buff[right] = buff[right], buff[left]
		left, right = left+1, right-1
	}
	return string(buff)
}
