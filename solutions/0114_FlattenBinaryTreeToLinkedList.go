package solutions

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	dummy := &TreeNode{}
	end := dummy
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		end.Left = nil
		end.Right = root
		end = root
		left, right := root.Left, root.Right
		if left != nil {
			traversal(left)
		}
		if right != nil {
			traversal(right)
		}
	}
	traversal(root)
}
