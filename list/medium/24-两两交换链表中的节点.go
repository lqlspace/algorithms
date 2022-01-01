package medium

// 递归法，时间复杂度O(N),空间复杂度O(N)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new := head.Next
	head.Next = swapPairs(new.Next)
	new.Next = head

	return new
}

// 迭代法，时间复杂度O(N),空间复杂度O(1)
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := new(ListNode)
	sentinel.Next = head
	pre := sentinel
	for pre.Next != nil && pre.Next.Next != nil {
		node := pre.Next.Next
		pre.Next.Next = node.Next
		node.Next = pre.Next
		pre.Next = node
		pre = pre.Next.Next
	}

	return sentinel.Next
}
