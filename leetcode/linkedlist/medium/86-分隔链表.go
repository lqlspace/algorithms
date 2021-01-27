package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间复杂度O(N)，空间复杂度O(1)
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	sentinel1 := &ListNode{Next:head}
	cur := sentinel1
	sentinel2 := new(ListNode)
	pre := sentinel2
	for cur.Next != nil {
		if cur.Next.Val < x {
			node := cur.Next
			cur.Next = cur.Next.Next
			node.Next = nil
			pre.Next = node
			pre = pre.Next
		} else {
			cur = cur.Next
		}
	}
	pre.Next = sentinel1.Next

	return sentinel2.Next
}
