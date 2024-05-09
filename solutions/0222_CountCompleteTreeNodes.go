package solutions

func countNodes(root *TreeNode) int {
	leftHeight := 0
	for node := root; nil != node; node = node.Left {
		leftHeight++
	}

	rightHeight := 0
	for node := root; nil != node; node = node.Right {
		rightHeight++
	}

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}
