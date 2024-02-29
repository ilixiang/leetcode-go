package solutions

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	tail := head
	for tail.Next != nil {
		node := tail.Next
		tail.Next = node.Next
		node.Next = dummy.Next
		dummy.Next = node
	}
	return dummy.Next
}
