package solutions

func GenerateTrees(n int) []*TreeNode {
	var generate func(int, int) []*TreeNode
	generate = func(left int, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}
		if left == right {
			return []*TreeNode{{Val: left, Left: nil, Right: nil}}
		}

		trees := make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			leftChildren := generate(left, i-1)
			rightChildren := generate(i+1, right)
			for _, leftChild := range leftChildren {
				for _, rightChild := range rightChildren {
					root := &TreeNode{Val: i, Left: leftChild, Right: rightChild}
					trees = append(trees, root)
				}
			}
		}
		return trees
	}
	return generate(1, n)
}
