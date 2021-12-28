package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 插入排序，时间复杂度O(N^2)，空间复杂度O（1）
func insertionSortList(head *ListNode) *ListNode {
	sentinel := new(ListNode)
	for head != nil {
		node := head
		head = head.Next
		pre := sentinel
		for pre.Next != nil && pre.Next.Val <= node.Val {
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
	}

	return sentinel.Next
}
