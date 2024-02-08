package solutions

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		} else if cur != prev.Next {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}
