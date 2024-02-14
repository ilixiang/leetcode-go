package solutions

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MinDepth(root.Left)
	right := MinDepth(root.Right)
	if left == 0 || right == 0 {
		return 1 + left + right
	} else {
		return 1 + min(left, right)
	}
}
