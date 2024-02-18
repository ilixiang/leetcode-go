package solutions

import "leetcode/structures"

type GraphNode = structures.GraphNode

func CloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	MAX_VAL := 100
	cache := make([]*GraphNode, MAX_VAL+1)

	var dfs func(*GraphNode) *GraphNode
	dfs = func(node *GraphNode) *GraphNode {
		nodeCopy := cache[node.Val]
		if nodeCopy != nil {
			return nodeCopy
		}

		nodeCopy = &GraphNode{Val: node.Val, Neighbors: make([]*GraphNode, 0, len(node.Neighbors))}
		cache[node.Val] = nodeCopy
		for _, neighbor := range node.Neighbors {
			neighborCopy := dfs(neighbor)
			nodeCopy.Neighbors = append(nodeCopy.Neighbors, neighborCopy)
		}
		return nodeCopy
	}
	return dfs(node)

}
