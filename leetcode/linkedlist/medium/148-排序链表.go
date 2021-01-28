package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 自顶向下归并排序，时间复杂度O(nlogn)，空间复杂度O(N)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	back := sortList(slow.Next)
	slow.Next = nil
	front := sortList(head)

	sentinel := new(ListNode)
	pre := sentinel
	for back != nil || front != nil {
		if front == nil {
			pre.Next = back
			break
		} else if back == nil {
			pre.Next = front
			break
		} else if front.Val < back.Val {
			pre.Next = front
			pre = pre.Next
			front = front.Next
		} else {
			pre.Next = back
			pre = pre.Next
			back = back.Next
		}
	}

	return sentinel.Next
}
