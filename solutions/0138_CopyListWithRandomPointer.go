package solutions

import "leetcode/structures"

type RandomPointerNode = structures.RandomPointerNode

func CopyRandomList(head *RandomPointerNode) *RandomPointerNode {
	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		nodeCopy := &RandomPointerNode{Val: node.Val, Next: node.Next}
		node.Next = nodeCopy
		node = nodeCopy.Next
	}

	node = head
	for node != nil {
		nodeCopy := node.Next
		if node.Random != nil {
			nodeCopy.Random = node.Random.Next
		}
		node = nodeCopy.Next
	}

	headCopy := head.Next
	node = head
	for node != nil {
		nodeCopy := node.Next
		node.Next = nodeCopy.Next
		if node.Next != nil {
			nodeCopy.Next = node.Next.Next
		}
		node = node.Next
	}
	return headCopy
}
