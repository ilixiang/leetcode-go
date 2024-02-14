package solutions

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	INT_MIN := -int(uint(1<<31-1)) - 1
	var recursive func(*TreeNode) (int, int)
	recursive = func(root *TreeNode) (int, int) {
		if root == nil {
			return INT_MIN, INT_MIN
		}

		val, left, right := root.Val, root.Left, root.Right
		leftSum, leftPathSum := recursive(left)
		rightSum, rightPathSum := recursive(right)

		rootSum := max(leftSum, rightSum, val, val+leftPathSum, val+rightPathSum, val+leftPathSum+rightPathSum)
		rootPathSum := max(val, val+leftPathSum, val+rightPathSum)
		return rootSum, rootPathSum
	}

	rootSum, _ := recursive(root)
	return rootSum
}
