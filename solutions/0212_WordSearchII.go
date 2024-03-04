package solutions

func FindWords(board [][]byte, words []string) []string {
	type TrieNode struct {
		nexts [26]*TrieNode
		end   bool
	}

	root := new(TrieNode)
	m, n := len(board), len(board[0])
	maxLen := m * n
	for _, word := range words {
		i, node := 0, root
		for ; i < maxLen && i < len(word); i++ {
			idx := int(word[i] - 'a')
			next := node.nexts[idx]
			if next == nil {
				next = new(TrieNode)
				node.nexts[idx] = next
			}
			node = next
		}
		node.end = i == len(word)
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := make([]string, 0, len(words))
	buf := make([]byte, 0, m*n)
	var dfs func(int, int, *TrieNode)
	dfs = func(row int, col int, node *TrieNode) {
		if row < 0 || row >= m || col < 0 || col >= n || visited[row][col] {
			return
		}

		visited[row][col] = true
		buf = append(buf, board[row][col])

		idx := int(board[row][col] - 'a')
		next := node.nexts[idx]
		if next != nil {
			if next.end {
				res = append(res, string(buf))
				next.end = false
			}
			dfs(row+1, col, next)
			dfs(row-1, col, next)
			dfs(row, col+1, next)
			dfs(row, col-1, next)
		}

		buf = buf[:len(buf)-1]
		visited[row][col] = false
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			dfs(row, col, root)
			if len(res) == len(words) {
				return res
			}
		}

	}
	return res
}
