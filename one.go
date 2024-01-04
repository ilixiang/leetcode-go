package one

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		diff := target - num
		j, ok := m[diff]
		if ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	n1, n2, tail := l1, l2, dummy
	carry := 0
	for nil != n1 && nil != n2 {
		cur := new(ListNode)
		cur.Val = (n1.Val + n2.Val + carry) % 10
		carry = (n1.Val + n2.Val + carry) / 10
		tail.Next = cur
		n1, n2, tail = n1.Next, n2.Next, cur
	}

	var n *ListNode
	if nil != n1 {
		n = n1
	} else {
		n = n2
	}

	for nil != n {
		cur := new(ListNode)
		cur.Val = (n.Val + carry) % 10
		carry = (n.Val + carry) / 10
		tail.Next = cur
		n, tail = n.Next, cur
	}

	if carry != 0 {
		cur := new(ListNode)
		cur.Val = carry
		tail.Next = cur
	}
	return dummy.Next
}
