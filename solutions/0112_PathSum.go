package solutions

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	} else {
		diff := targetSum - root.Val
		return HasPathSum(root.Left, diff) || HasPathSum(root.Right, diff)
	}
}
