/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{0, nil}

	var carryover, number int
	l3 := head

	for l, m := l1, l2; l != nil || m != nil; {
		if l == nil && m != nil {
			number = 0 + m.Val + carryover

		} else if l != nil && m == nil {
			number = l.Val + 0 + carryover
		} else {
			number = l.Val + m.Val + carryover
		}

		if number >= 10 {
			carryover = number / 10
			number = number % 10
		} else {

			carryover = 0
		}
		l3.Val = number

		if m != nil {
			m = m.Next
		}
		if l != nil {
			l = l.Next
		}
		if l != nil || m != nil {
			newNode := &ListNode{0, nil}
			l3.Next = newNode
			l3 = l3.Next
		}

	}
	if carryover > 0 {
		newNode := &ListNode{0, nil}
		l3.Next = newNode
		l3 = l3.Next
		l3.Val = carryover
	}

	return head

}
