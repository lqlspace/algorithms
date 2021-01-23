package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归法，时间复杂度O(N),空间复杂度O(N)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := swapPairs(head.Next.Next)
	head.Next.Next = nil
	head = reverseList(head)
	head.Next.Next = p

	return head
}


func reverseList(head *ListNode) *ListNode {
	sentinel := new(ListNode)
	for head != nil {
		p := head
		head = head.Next
		p.Next = sentinel.Next
		sentinel.Next = p
	}

	return sentinel.Next
}

// 迭代法，时间复杂度O(N),空间复杂度O(1)
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := new(ListNode)
	sentinel.Next = head
	cur := sentinel
	for cur.Next != nil && cur.Next.Next != nil {
		skip := cur.Next.Next.Next
		cur.Next.Next.Next = nil
		cur.Next = reverseList(cur.Next)
		cur.Next.Next.Next = skip
		cur = cur.Next.Next
	}

	return sentinel.Next
}


