package solutions

func ZigzagConversion(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}
	runes := []rune(s)
	buf := make([]rune, 0, len(s))
	n := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		j := 0
		for j*n+i < len(s) {
			buf = append(buf, runes[j*n+i])
			if 0 < i && i < numRows-1 && (j+1)*n-i < len(s) {
				buf = append(buf, runes[(j+1)*n-i])
			}
			j += 1
		}
	}
	return string(buf)
}
