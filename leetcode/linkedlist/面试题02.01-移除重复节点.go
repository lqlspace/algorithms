package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 借助hash表（时间复杂度O(N), 空间复杂度O(N)）
func removeDuplicateNodesM1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[int]struct{})
	m[head.Val] = struct{}{}
	cur := head
	for cur.Next != nil {
		if _, ok := m[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			m[cur.Next.Val] = struct{}{}
			cur = cur.Next
		}
	}

	return head
}

// 内外双循环（时间复杂度O(N^2)，空间复杂度O(1)）
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	outer := head
	for outer != nil {
		inner := outer
		for inner.Next != nil {
			if inner.Next.Val == outer.Val {
				inner.Next = inner.Next.Next
			} else {
				inner = inner.Next
			}
		}
		outer = outer.Next
	}

	return head
}
