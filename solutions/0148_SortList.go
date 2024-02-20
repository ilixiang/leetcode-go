package solutions

func SortList(head *ListNode) *ListNode {
	merge := func(l1 *ListNode, l2 *ListNode) *ListNode {
		dummy := &ListNode{}
		n, n1, n2 := dummy, l1, l2
		for n1 != nil && n2 != nil {
			if n1.Val < n2.Val {
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

	var sort func(*ListNode) *ListNode
	sort = func(l *ListNode) *ListNode {
		if l == nil || l.Next == nil {
			return l
		}

		dummy := &ListNode{Next: l}
		s, f := dummy, dummy
		for f != nil && f.Next != nil {
			s = s.Next
			f = f.Next.Next
		}
		l1, l2 := l, s.Next
		s.Next = nil
		l1 = sort(l1)
		l2 = sort(l2)
		return merge(l1, l2)
	}
	return sort(head)
}
