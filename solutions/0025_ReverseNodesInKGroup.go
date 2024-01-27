package solutions

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	n1, n2 := dummy, head
	counter := 1
	for n2 != nil {
		for counter = 1; n2.Next != nil && counter < k; counter++ {
			tmp := n1.Next
			n1.Next = n2.Next
			n2.Next = n2.Next.Next
			n1.Next.Next = tmp
		}
		if counter < k {
			break
		} else {
			n1 = n2
			n2 = n2.Next
		}
	}
	if counter < k {
		n2 = n1.Next
		for n2.Next != nil {
			tmp := n1.Next
			n1.Next = n2.Next
			n2.Next = n2.Next.Next
			n1.Next.Next = tmp
		}
	}

	return dummy.Next
}
