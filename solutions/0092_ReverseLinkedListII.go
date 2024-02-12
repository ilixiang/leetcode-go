package solutions

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	n1 := dummy
	for i := 1; i < left; i++ {
		n1 = n1.Next
	}

	count := right - left
	n2 := n1.Next
	if count <= 0 || n2 == nil {
		return head
	}

	for n2.Next != nil && count > 0 {
		n := n2.Next
		n2.Next = n.Next
		n.Next = n1.Next
		n1.Next = n
		count--
	}
	return dummy.Next
}
