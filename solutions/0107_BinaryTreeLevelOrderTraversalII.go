package solutions

func LevelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		buf := make([]int, 0)
		count := len(queue)
		for count != 0 {
			node := queue[0]
			buf = append(buf, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			count--
		}
		result = append(result, buf)
	}

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
