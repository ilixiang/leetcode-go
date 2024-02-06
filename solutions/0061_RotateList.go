package solutions

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}

	total := 0
	node := dummy
	for node.Next != nil {
		total++
		node = node.Next
	}
	tail := node

	k %= total
	if k == 0 {
		return head
	}

	node = dummy
	for i := 0; i < total-k; i++ {
		node = node.Next
	}

	tail.Next = dummy.Next
	dummy.Next = node.Next
	node.Next = nil
	return dummy.Next
}
