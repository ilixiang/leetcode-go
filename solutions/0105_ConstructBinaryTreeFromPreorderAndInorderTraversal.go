package solutions

func BuildTreeFromPreorderAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	inorderMap := make([]int, 6001)
	for i := 0; i < len(inorderMap); i++ {
		inorderMap[i] = -1
	}
	for i, num := range inorder {
		inorderMap[num+3000] = i
	}

	var buildRecursive func(int, int, int) *TreeNode
	buildRecursive = func(i1 int, i2 int, size int) *TreeNode {
		if size == 0 {
			return nil
		}

		val := preorder[i1]
		root := &TreeNode{Val: val}
		inorderIdx := inorderMap[val+3000]
		leftSize, rightSize := inorderIdx-i2, i2+size-inorderIdx-1
		root.Left = buildRecursive(i1+1, i2, leftSize)
		root.Right = buildRecursive(i1+1+leftSize, inorderIdx+1, rightSize)
		return root
	}
	return buildRecursive(0, 0, len(inorder))
}
