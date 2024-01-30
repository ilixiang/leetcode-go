package solutions

func findSubstring(s string, words []string) []int {
	rev := make([]int, 0, 10)
	wordLen := len(words[0])
	wordMap := map[string]int{}
	substringLen := wordLen * len(words)
	for _, word := range words {
		wordMap[word] += 1
	}

	for idx := 0; idx < wordLen; idx++ {
		start := idx
		end := idx + wordLen
		for ; end <= len(s); end += wordLen {
			word := s[end-wordLen : end]
			remain, existed := wordMap[word]
			if existed && remain > 0 {
				wordMap[word] = remain - 1
				if end-start == substringLen {
					rev = append(rev, start)
					passingWord := s[start : start+wordLen]
					wordMap[passingWord] += 1
					start += wordLen
				}
			} else if existed {
				// restore wordMap util encounter word
				passingWord := s[start : start+wordLen]
				for passingWord != word {
					wordMap[passingWord] += 1
					start += wordLen
					passingWord = s[start : start+wordLen]
				}
				start += wordLen
			} else {
				// restore wordMap
				for ; start < end-wordLen; start += wordLen {
					passingWord := s[start : start+wordLen]
					wordMap[passingWord] += 1
				}
				start = end
			}
		}

		for ; start < end-wordLen; start += wordLen {
			wordMap[s[start:start+wordLen]] += 1
		}
	}

	return rev
}
