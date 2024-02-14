package solutions

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var buildRecursive func(*ListNode, *ListNode) *TreeNode
	buildRecursive = func(left *ListNode, right *ListNode) *TreeNode {
		if left == right {
			return nil
		}

		slow, fast := left, left
		for fast != right && fast.Next != right {
			slow = slow.Next
			fast = fast.Next.Next
		}

		root := &TreeNode{Val: slow.Val, Left: buildRecursive(left, slow), Right: buildRecursive(slow.Next, right)}
		return root
	}
	return buildRecursive(head, nil)
}
