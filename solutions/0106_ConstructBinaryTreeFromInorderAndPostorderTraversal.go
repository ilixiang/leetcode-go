package solutions

func BuildTreeFromInorderAndPostorder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
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

		val := postorder[i2]
		root := &TreeNode{Val: val}
		inorderIdx := inorderMap[val+3000]
		leftSize, rightSize := inorderIdx-i1, i1+size-inorderIdx-1
		root.Left = buildRecursive(i1, i2-1-rightSize, leftSize)
		root.Right = buildRecursive(inorderIdx+1, i2-1, rightSize)
		return root
	}
	return buildRecursive(0, len(inorder)-1, len(inorder))
}
