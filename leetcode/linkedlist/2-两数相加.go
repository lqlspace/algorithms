package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间复杂度O(MAX(M, N))， 空间复杂度O（MAX(M, N)）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	last := sentinel

	var carry int
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		last.Next = node
		last = last.Next
	}

	if carry > 0 {
		node := new(ListNode)
		node.Val = carry
		last.Next = node
		last = last.Next
	}

	return sentinel.Next
}
