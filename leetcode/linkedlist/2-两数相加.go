package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	last := sentinel
	var ans int
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + ans
		ans = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		last.Next = node
		last = last.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + ans
		ans = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		last.Next = node
		last = last.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + ans
		ans = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		last.Next = node
		last = last.Next
		l2 = l2.Next
	}

	if ans > 0 {
		node := new(ListNode)
		node.Val = ans
		last.Next = node
		last = last.Next
	}

	return sentinel.Next
}
