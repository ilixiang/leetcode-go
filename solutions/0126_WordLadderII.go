package solutions

type WordLadderNode struct {
	Word  string
	Prevs []*WordLadderNode
}

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordLen := len(wordList[0])
	wordMap := map[string][]string{}
	for _, word := range wordList {
		for i := 0; i < wordLen; i++ {
			wildcardWord := word[:i] + "." + word[i+1:]
			wordMap[wildcardWord] = append(wordMap[wildcardWord], word)
		}
	}

	// BFS
	visitedNodeMap := map[string]*WordLadderNode{}
	queue := make([]*WordLadderNode, 1)
	beginWordNode := new(WordLadderNode)
	beginWordNode.Word = beginWord
	beginWordNode.Prevs = append(beginWordNode.Prevs, nil)
	queue[0] = beginWordNode
	for len(queue) != 0 && visitedNodeMap[endWord] == nil {
		createdNodeMap := map[string]*WordLadderNode{}
		for count := len(queue); count > 0; count-- {
			wordNode := queue[0]
			word := wordNode.Word
			queue = queue[1:]
			for i := 0; i < wordLen; i++ {
				wildcardWord := word[:i] + "." + word[i+1:]
				nextWords := wordMap[wildcardWord]
				for _, nextWord := range nextWords {
					if visitedNodeMap[nextWord] != nil {
						continue
					}
					createdNode := createdNodeMap[nextWord]
					if createdNode == nil {
						createdNode = new(WordLadderNode)
						createdNode.Word = nextWord
						createdNodeMap[nextWord] = createdNode
						queue = append(queue, createdNode)
					}
					createdNode.Prevs = append(createdNode.Prevs, wordNode)
				}
			}
		}

		for createdNodeWord, node := range createdNodeMap {
			visitedNodeMap[createdNodeWord] = node
		}
	}

	// DFS
	results := make([][]string, 0)
	node := visitedNodeMap[endWord]
	if node == nil {
		return results
	}
	buf := make([]string, 0)
	var dfs func(*WordLadderNode)
	dfs = func(node *WordLadderNode) {
		if node == nil {
			result := make([]string, len(buf))
			copy(result, buf)
			left, right := 0, len(result)-1
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			results = append(results, result)
			return
		}

		buf = append(buf, node.Word)
		for _, prev := range node.Prevs {
			dfs(prev)
		}
		buf = buf[:len(buf)-1]
	}
	dfs(node)
	return results
}
