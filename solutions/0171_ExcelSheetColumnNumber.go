package solutions

func TitleToNumber(columnTitle string) int {
	res := 0
	for _, letter := range columnTitle {
		res = res*26 + int(letter-'A') + 1
	}
	return res
}
