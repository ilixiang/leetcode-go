package solutions

func SumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}

	var recursive func(*TreeNode, int)
	recursive = func(root *TreeNode, prev int) {
		sum := root.Val + prev*10
		if root.Left == nil && root.Right == nil {
			res += sum
			return
		}

		if root.Left != nil {
			recursive(root.Left, sum)
		}
		if root.Right != nil {
			recursive(root.Right, sum)
		}
	}
	recursive(root, 0)
	return res
}
