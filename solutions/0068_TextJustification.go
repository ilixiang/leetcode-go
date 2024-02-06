package solutions

import (
	"strings"
)

func FullJustifiy(words []string, maxWidth int) []string {
	result := make([]string, 0, len(words))
	sb := strings.Builder{}
	selected := make([]string, 0, len(words))
	i, width := 0, 0
	for i < len(words) {
		word := words[i]
		occupation := 0
		if width == 0 {
			occupation = len(word)
		} else {
			occupation = 1 + len(word)
		}

		if width+occupation <= maxWidth {
			selected = append(selected, word)
			width += occupation
			i++
		} else {
			num := len(selected)
			space := maxWidth - width + num - 1
			if num == 1 {
				sb.WriteString(selected[0])
				for ; space > 0; space-- {
					sb.WriteByte(' ')
				}
			} else {
				avgSpace := space / (num - 1)
				extraSpace := space % (num - 1)
				for i := 0; i < num-1; i++ {
					sb.WriteString(selected[i])
					for j := 0; j < avgSpace; j++ {
						sb.WriteByte(' ')
					}
					if extraSpace != 0 {
						sb.WriteByte(' ')
						extraSpace--
					}
				}
				sb.WriteString(selected[len(selected)-1])
			}

			selected = selected[:0]
			result = append(result, sb.String())
			sb.Reset()
			width = 0
		}
	}

	for i := 0; i < len(selected)-1; i++ {
		sb.WriteString(selected[i])
		sb.WriteByte(' ')
	}
	sb.WriteString(selected[len(selected)-1])
	for i := maxWidth - width; i > 0; i-- {
		sb.WriteByte(' ')
	}
	result = append(result, sb.String())
	return result
}
