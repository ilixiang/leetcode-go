package solutions

func PathSum(root *TreeNode, targetSum int) [][]int {
	var recursive func(*TreeNode, int) [][]int
	recursive = func(root *TreeNode, target int) [][]int {
		results := make([][]int, 0)
		if root == nil {
			return results
		}

		val, left, right := root.Val, root.Left, root.Right
		if left == nil && right == nil {
			if val == target {
				results = append(results, []int{target})
			}
			return results
		}

		leftResults := recursive(left, target-val)
		rightResults := recursive(right, target-val)
		results = append(results, leftResults...)
		results = append(results, rightResults...)
		for i := range results {
			results[i] = append(results[i], val)
		}
		return results
	}

	results := recursive(root, targetSum)
	for _, reversedResult := range results {
		left, right := 0, len(reversedResult)-1
		for left < right {
			reversedResult[left], reversedResult[right] = reversedResult[right], reversedResult[left]
			left++
			right--
		}
	}
	return results
}
