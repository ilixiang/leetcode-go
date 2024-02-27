package solutions

func RightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		res = append(res, node.Val)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return res
}
