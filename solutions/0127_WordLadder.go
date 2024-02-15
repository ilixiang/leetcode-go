package solutions

func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordLen := len(wordList[0])
	wordMap := map[string][]string{}
	for _, word := range wordList {
		for i := 0; i < wordLen; i++ {
			wildcardWord := word[:i] + "." + word[i+1:]
			wordMap[wildcardWord] = append(wordMap[wildcardWord], word)
		}
	}

	// BFS
	res := 1
	visitedNodeMap := map[string]*struct{}{}
	queue := make([]string, 1)
	queue[0] = beginWord
	for len(queue) != 0 && visitedNodeMap[endWord] == nil {
		for count := len(queue); count > 0; count-- {
			word := queue[0]
			queue = queue[1:]
			for i := 0; i < wordLen; i++ {
				wildcardWord := word[:i] + "." + word[i+1:]
				nextWords := wordMap[wildcardWord]
				for _, nextWord := range nextWords {
					if visitedNodeMap[nextWord] == nil {
						visitedNodeMap[nextWord] = new(struct{})
						queue = append(queue, nextWord)
					}
				}
			}
		}
		res++
	}

	if visitedNodeMap[endWord] == nil {
		return 0
	}
	return res
}
