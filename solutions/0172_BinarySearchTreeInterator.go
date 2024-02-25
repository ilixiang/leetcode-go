package solutions

type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root, stack: []*TreeNode{}}
}

func (iterator *BSTIterator) Next() int {
	node := iterator.cur
	for node != nil {
		iterator.stack = append(iterator.stack, node)
		node = node.Left
	}
	node = iterator.stack[len(iterator.stack)-1]
	iterator.stack = iterator.stack[:len(iterator.stack)-1]
	iterator.cur = node.Right
	return node.Val
}

func (iterator *BSTIterator) HasNext() bool {
	return len(iterator.stack) != 0 || iterator.cur != nil
}
