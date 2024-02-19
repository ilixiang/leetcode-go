package solutions

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	n1, n2 := head, head.Next
	for n2 != nil && n2.Next != nil && n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	return n1 == n2
}
