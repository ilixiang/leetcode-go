package solutions

func SortedArrayToBST(nums []int) *TreeNode {
	var buildRecursive func(int, int) *TreeNode
	buildRecursive = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := (left + right) >> 1
		root := TreeNode{Val: nums[mid], Left: buildRecursive(left, mid-1), Right: buildRecursive(mid+1, right)}
		return &root
	}
	return buildRecursive(0, len(nums)-1)
}
