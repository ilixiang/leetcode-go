package solutions

func Partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}
	n1, n2 := l1, l2
	n := head
	for n != nil {
		if n.Val < x {
			n1.Next = n
			n1 = n
		} else {
			n2.Next = n
			n2 = n
		}
		n = n.Next
	}
	n1.Next = l2.Next
	n2.Next = nil
	return l1.Next
}
