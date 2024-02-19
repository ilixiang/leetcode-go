package solutions

func ReorderList(head *ListNode) {
	l1 := &ListNode{}
	l2 := &ListNode{}

	n1, n2 := head, head
	for n2 != nil && n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	if n1 == n2 || n1.Next == nil {
		return
	}

	l1.Next = head
	n2 = n1.Next
	n1.Next = nil

	l2.Next = n2
	for n2.Next != nil {
		node := n2.Next
		n2.Next = node.Next
		node.Next = l2.Next
		l2.Next = node
	}

	n1 = l1.Next
	n2 = l2.Next
	for n2 != nil {
		tmp := n2.Next
		n2.Next = n1.Next
		n1.Next = n2
		n1 = n2.Next
		n2 = tmp
	}
}
