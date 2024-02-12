package solutions

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(*TreeNode, *TreeNode) bool
	isMirror = func(r1 *TreeNode, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		} else if r1 == nil || r2 == nil {
			return false
		} else {
			return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
		}
	}
	return isMirror(root.Left, root.Right)
}
