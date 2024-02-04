package solutions

func MergeKLists(lists []*ListNode) *ListNode {
	var merge func([]*ListNode, []*ListNode) *ListNode
	merge = func(ls1 []*ListNode, ls2 []*ListNode) *ListNode {
		var l1, l2 *ListNode = nil, nil
		if len(ls1) == 1 {
			l1 = ls1[0]
		} else if len(ls1) > 1 {
			l1 = merge(ls1[:len(ls1)/2], ls1[len(ls1)/2:])
		}
		if len(ls2) == 1 {
			l2 = ls2[0]
		} else if len(ls2) > 1 {
			l2 = merge(ls2[:len(ls2)/2], ls2[len(ls2)/2:])
		}

		n1, n2 := l1, l2
		dummy := &ListNode{Val: 0, Next: nil}
		n := dummy
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
		if n1 == nil {
			n.Next = n2
		} else {
			n.Next = n1
		}
		return dummy.Next
	}
	return merge(lists[:len(lists)/2], lists[len(lists)/2:])
}
