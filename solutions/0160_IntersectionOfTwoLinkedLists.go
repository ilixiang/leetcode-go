package solutions

func GetIntersectionNode(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := l1, l2
	for n1 != nil && n2 != nil {
		n1 = n1.Next
		n2 = n2.Next
	}

	if n1 == nil {
		n1 = l2
	}
	if n2 == nil {
		n2 = l1
	}
	for n1 != nil && n2 != nil {
		n1 = n1.Next
		n2 = n2.Next
	}

	if n1 == nil {
		n1 = l2
	}
	if n2 == nil {
		n2 = l1
	}
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}
