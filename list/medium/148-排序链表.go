package medium


// 时间复杂度O(N^2)，会超时
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	for head != nil {
		node := head
		head = head.Next
		pre := dummy
		for pre.Next != nil && pre.Next.Val < node.Val{
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
	}

	return dummy.Next
}

// 自顶向下归并排序，时间复杂度O(nlogn)，空间复杂度O(N)
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := cutTwoHalf(head)
	return merge(sortList3(left), sortList3(right))
}

func cutTwoHalf(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	right := slow.Next
	slow.Next = nil
	return head, right
}

func merge(left, right *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}

	if left != nil {
		pre.Next = left
	}
	if right != nil {
		pre.Next = right
	}

	return dummy.Next
}
