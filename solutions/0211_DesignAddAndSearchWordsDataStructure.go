package solutions

type WordDictionary struct {
	nexts [26]*WordDictionary
	end   bool
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{}
}

func (dictionary *WordDictionary) AddWord(word string) {
	node := dictionary
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		next := node.nexts[idx]
		if next == nil {
			next = new(WordDictionary)
			node.nexts[idx] = next
		}
		node = next
	}
	node.end = true
}

func (dictionary *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return dictionary != nil && dictionary.end
	}
	if dictionary == nil {
		return false
	}

	idx := int(word[0] - 'a')
	if word[0] == '.' {
		for _, next := range dictionary.nexts {
			if next.Search(word[1:]) {
				return true
			}
		}
		return false
	} else {
		return dictionary.nexts[idx].Search(word[1:])
	}
}
