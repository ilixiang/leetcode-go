package solutions

func IsBalanced(root *TreeNode) bool {
	var recursive func(*TreeNode) (bool, int)
	recursive = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}

		leftBalanced, leftHeight := recursive(root.Left)
		if leftBalanced {
			rightBalanced, rightHeight := recursive(root.Right)
			heightDiff := leftHeight - rightHeight
			balanced := leftBalanced && rightBalanced && -1 <= heightDiff && heightDiff <= 1
			return balanced, 1 + max(leftHeight, rightHeight)
		}
		return false, 0
	}
	balanced, _ := recursive(root)
	return balanced
}
