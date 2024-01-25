package solutions

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	n1, n2 := dummy, dummy

	i := 0
	for i = 0; i <= n && nil != n2; i++ {
		n2 = n2.Next
	}
	if i < n {
		return head
	}

	for nil != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	if n1 != nil && n1.Next != nil {
		n1.Next = n1.Next.Next
	}
	return dummy.Next
}
