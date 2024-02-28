package solutions

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil {
		node := prev.Next
		if node.Val == val {
			prev.Next = node.Next
		} else {
			prev = node
		}
	}
	return dummy.Next
}
