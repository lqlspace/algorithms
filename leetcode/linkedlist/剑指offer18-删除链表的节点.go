package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodeMethod1(head *ListNode, val int) *ListNode {
	for head == nil {
		return nil
	}

	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			head.Next = nil
			break
		}
		pre = head
		head = head.Next
	}

	return sentinel.Next
}
