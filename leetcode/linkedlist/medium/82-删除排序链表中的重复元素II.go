package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代法
const min  = -65534
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := &ListNode{Next:head}
	var dup = min
	cur := sentinel
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			dup = cur.Next.Val
			cur.Next = cur.Next.Next.Next
		} else {
			if cur.Next.Val == dup {
				cur.Next = cur.Next.Next
			} else {
				dup = min
				cur = cur.Next
			}
		}
	}

	if cur.Next != nil && cur.Next.Val == dup {
		cur.Next = cur.Next.Next
	}

	return sentinel.Next
}

// 递归法
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dup bool
	var recurFunc func(cur *ListNode) *ListNode
	recurFunc = func(cur *ListNode) *ListNode {
		if cur.Next == nil {
			return cur
		}

		next := recurFunc(cur.Next)
		var ret *ListNode
		switch {
		case cur.Val == next.Val:
			dup = true
			ret = next
		case cur.Val != next.Val && !dup:
			ret = cur
		case cur.Val != next.Val && dup:
			dup = false
			cur.Next = next.Next
			ret = cur
		}
		return ret
	}

	head = recurFunc(head)
	if dup {
		head = head.Next
	}

	return head
}
