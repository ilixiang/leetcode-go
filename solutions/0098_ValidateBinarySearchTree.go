package solutions

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var recursive func(*TreeNode) (bool, int, int)
	recursive = func(root *TreeNode) (bool, int, int) {
		val, left, right := root.Val, root.Left, root.Right
		if left == nil && right == nil {
			return true, root.Val, root.Val
		} else if left == nil {
			rightBST, rightMin, rightMax := recursive(right)
			return val < rightMin && rightBST, val, rightMax
		} else if right == nil {
			leftBST, leftMin, leftMax := recursive(left)
			return val > leftMax && leftBST, leftMin, val
		} else {
			leftBST, leftMin, leftMax := recursive(left)
			rightBST, rightMin, rightMax := recursive(right)
			return val > leftMax && val < rightMin && leftBST && rightBST, leftMin, rightMax
		}
	}

	rootBTS, _, _ := recursive(root)
	return rootBTS
}
