package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1, p2 := head, head.Next
	for p2 != nil {
		if p2.Val != p1.Val {
			p1.Next = p2
			p1 = p1.Next
		}
		p2 = p2.Next
	}
	p1.Next = p2

	return head
}
