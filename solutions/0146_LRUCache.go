package solutions

type CacheNode struct {
	Key  int
	Val  int
	Prev *CacheNode
	Next *CacheNode
}

type LRUCache struct {
	capacity int
	storage  map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
}

func Constructor(capacity int) LRUCache {
	storage := make(map[int]*CacheNode, capacity)
	head := new(CacheNode)
	tail := new(CacheNode)
	head.Next = tail
	tail.Prev = head
	return LRUCache{capacity: capacity, storage: storage, head: head, tail: tail}
}

func (cache *LRUCache) Get(key int) int {
	node := cache.storage[key]
	if node == nil {
		return -1
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = cache.head.Next
	node.Prev = cache.head
	cache.head.Next = node
	node.Next.Prev = node
	return node.Val
}

func (cache *LRUCache) Put(key int, val int) {
	node := cache.storage[key]
	if node != nil {
		node.Val = val
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	} else if len(cache.storage) == cache.capacity {
		eviction := cache.tail.Prev
		eviction.Prev.Next = eviction.Next
		eviction.Next.Prev = eviction.Prev
		delete(cache.storage, eviction.Key)
		node = &CacheNode{Key: key, Val: val}
		cache.storage[key] = node
	} else {
		node = &CacheNode{Key: key, Val: val}
		cache.storage[key] = node
	}

	node.Next = cache.head.Next
	node.Prev = cache.head
	cache.head.Next = node
	node.Next.Prev = node
}
