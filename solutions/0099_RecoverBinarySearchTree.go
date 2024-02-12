package solutions

func RecoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0, 1000)

	node := root
	var prev *TreeNode = nil
	var first, last *TreeNode = nil, nil
	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = popped.Right
			if prev != nil && prev.Val > popped.Val {
				if first == nil {
					first = prev
				}
				last = popped
			}
			prev = popped
		}
	}
	first.Val, last.Val = last.Val, first.Val
}
