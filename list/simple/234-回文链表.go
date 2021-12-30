package simple

// 数组法（时间复杂度O(N),空间复杂度O(N)）
func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return true
	}
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	start, end := 0, len(arr)-1
	for start <= end {
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// 递归法(时间复杂度O(N), 空间复杂度O(N))
func isPalindrome2(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(node *ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}

	return recursivelyCheck(head)
}

// 快慢指针+反转链表（会修改链表结构，非并发安全，时间复杂度O(N)，空间复杂度O(1)）
func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}

	halfNode := getHalfNode(head)
	startOfSecondHalf := reverseList(halfNode.Next)

	second := startOfSecondHalf
	for head != nil && second != nil {
		if head.Val != second.Val {
			return false
		}
		head = head.Next
		second = second.Next
	}

	// 恢复反转的链表
	halfNode.Next = reverseList(startOfSecondHalf)

	return true
}

func getHalfNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
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
