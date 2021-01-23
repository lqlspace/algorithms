package simple

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 哨兵节点（哨兵节点广泛应用于树和链表中，如伪头、伪尾、标记等，
//它们是纯功能的，通常不保存任何数据，主要目的是使链表标准化，
//如链表永不为空、永不无头、简化插入和删除）
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	sentinel := new(ListNode)
	pre, cur := sentinel, head
	for cur != nil {
		if cur.Val != val {
			pre.Next = cur
			pre = pre.Next
		}
		cur = cur.Next
	}
	pre.Next = nil

	return sentinel.Next
}
