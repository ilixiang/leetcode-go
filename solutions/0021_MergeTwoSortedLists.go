package solutions

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	n := dummy
	n1, n2 := list1, list2
	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}
		n = n.Next
	}
	if n1 != nil {
		n.Next = n1
	} else {
		n.Next = n2
	}
	return dummy.Next
}
