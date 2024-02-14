package solutions

import "leetcode/structures"

type Node = structures.Node

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummy := &Node{}
	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) != 0 {
		prev := dummy
		for count := len(queue); count > 0; count-- {
			node := queue[0]
			queue = queue[1:]
			prev.Next = node
			prev = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		prev.Next = nil
	}
	return root
}
