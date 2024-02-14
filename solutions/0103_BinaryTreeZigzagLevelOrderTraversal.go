package solutions

func ZigzagLevelOrder(root *TreeNode) [][]int {
	results := make([][]int, 0)
	if root == nil {
		return results
	}

	s1 := make([]*TreeNode, 1)
	s2 := make([]*TreeNode, 0)

	s1[0] = root
	for len(s1) != 0 {
		result := make([]int, 0)
		level := len(results)
		for count := len(s1); count > 0; count-- {
			node := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			result = append(result, node.Val)
			if level&1 == 1 {
				if node.Right != nil {
					s2 = append(s2, node.Right)
				}
				if node.Left != nil {
					s2 = append(s2, node.Left)
				}
			} else {
				if node.Left != nil {
					s2 = append(s2, node.Left)
				}
				if node.Right != nil {
					s2 = append(s2, node.Right)
				}
			}
		}
		results = append(results, result)
		s1, s2 = s2, s1
	}
	return results
}
