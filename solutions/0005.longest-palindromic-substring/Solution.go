package solutions

import "strings"

func longestPalindromicSubstring(s string) string {
	if len(s) <= 1 {
		return s
	}

	ss := make([]rune, 2*len(s)+1)
	for i, r := range s {
		ss[2*i] = '#'
		ss[2*i+1] = r
	}
	ss[2*len(s)] = '#'

	radiuses := make([]int, len(ss))

	maxCoverCenter := -1
	maxCoverRadius := 0

	for i := range ss {
		maxCoverIdx := maxCoverCenter + maxCoverRadius
		if i < maxCoverIdx {
			symmetry := 2*maxCoverCenter - i
			symmetryRadius := radiuses[symmetry]

			if i+symmetryRadius < maxCoverIdx {
				radiuses[i] = symmetryRadius
			} else if i+symmetryRadius == maxCoverIdx {
				left, right := 2*i-maxCoverIdx-1, maxCoverIdx+1
				for left >= 0 && right < len(ss) && ss[left] == ss[right] {
					left -= 1
					right += 1
				}
				maxCoverCenter, maxCoverRadius = i, right-i-1
				radiuses[i] = maxCoverRadius
			} else {
				radiuses[i] = maxCoverIdx - i
			}

		} else {
			left, right := i-1, i+1
			for left >= 0 && right < len(ss) && ss[left] == ss[right] {
				left -= 1
				right += 1
			}
			maxCoverCenter = i
			maxCoverRadius = right - i - 1
			radiuses[i] = right - i - 1
		}
	}

	maxRadiusCenter, maxRadius := 0, 0
	for i, radius := range radiuses {
		if radius >= maxRadius {
			maxRadiusCenter = i
			maxRadius = radius
		}
	}

	sb := strings.Builder{}
	left, right := maxRadiusCenter-maxRadius, maxRadiusCenter+maxRadius
	for i := left; i <= right; i++ {
		if (i-left)&1 == 1 {
			sb.WriteRune(ss[i])
		}
	}
	return sb.String()
}
