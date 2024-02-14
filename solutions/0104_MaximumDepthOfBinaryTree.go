package solutions

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}
