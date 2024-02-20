package solutions

func InsertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for node.Next != nil {
		prev := dummy
		for prev != node && prev.Next.Val < node.Next.Val {
			prev = prev.Next
		}

		if prev == node {
			node = node.Next
		} else {
			tmp := node.Next
			node.Next = tmp.Next
			tmp.Next = prev.Next
			prev.Next = tmp
		}
	}
	return dummy.Next
}
