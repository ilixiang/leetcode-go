package solutions

type Trie struct {
	nexts  [26]*Trie
	prefix int
	end    bool
}

func ConstructTrie() Trie {
	return Trie{}
}

func (root *Trie) Insert(word string) {
	node := root
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		next := node.nexts[idx]
		if next == nil {
			next = new(Trie)
			node.nexts[idx] = next
		}
		node.prefix++
		node = next
	}
	node.end = true
}

func (root *Trie) Search(word string) bool {
	i, node := 0, root
	for i < len(word) && node != nil {
		idx := int(word[i] - 'a')
		next := node.nexts[idx]
		node = next
		i++
	}
	return i == len(word) && node != nil && node.end
}

func (root *Trie) StartsWith(prefix string) bool {
	i, node := 0, root
	for i < len(prefix) && node != nil {
		idx := int(prefix[i] - 'a')
		next := node.nexts[idx]
		node = next
		i++
	}
	return i == len(prefix) && node != nil && (node.prefix != 0 || node.end)
}
