package solutions

func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	n1, n2 := head.Next, head.Next.Next
	for n2 != nil && n2.Next != nil && n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next.Next
	}

	if n1 != n2 {
		return nil
	}

	n2 = head
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}
