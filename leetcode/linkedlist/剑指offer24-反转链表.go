package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归法
func reverseListM1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := reverseListM1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return pre
}

// 迭代法
func reverseListM2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentinel := new(ListNode)
	for head != nil {
		p := head
		head = head.Next
		p.Next = sentinel.Next
		sentinel.Next = p
	}

	return sentinel.Next
}
