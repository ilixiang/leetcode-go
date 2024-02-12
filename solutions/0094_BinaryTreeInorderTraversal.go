package solutions

import "leetcode/structures"

type TreeNode = structures.TreeNode

func InorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, popped.Val)
			node = popped.Right
		}
	}
	return result
}
