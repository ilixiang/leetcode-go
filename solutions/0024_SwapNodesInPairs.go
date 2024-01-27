package solutions

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	n1, n2 := dummy, head
	for n1 != nil && n2 != nil && n2.Next != nil {
		n1.Next = n2.Next
		n2.Next = n2.Next.Next
		n1.Next.Next = n2
		n1 = n2
		n2 = n1.Next
	}
	return dummy.Next
}
