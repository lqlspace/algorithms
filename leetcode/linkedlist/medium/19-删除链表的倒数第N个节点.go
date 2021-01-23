package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两次遍历
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var num int
	cur := head
	for cur != nil {
		num++
		cur = cur.Next
	}

	if num < n {
		return nil
	}

	sentinel := new(ListNode)
	sentinel.Next = head

	pos := num - n
	cur = sentinel
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return sentinel.Next
}
