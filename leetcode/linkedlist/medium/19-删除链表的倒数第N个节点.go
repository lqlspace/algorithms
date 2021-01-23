package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两次遍历(时间复杂度O(N)，空间复杂度O(1))
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var num int
	cur := head
	for cur != nil {
		num++
		cur = cur.Next
	}

	if num < n {
		return nil
	}

	sentinel := new(ListNode)
	sentinel.Next = head

	pos := num - n
	cur = sentinel
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return sentinel.Next
}

// 递归，时间复杂度O(N)，空间复杂度O(N)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var num int
	var checkRecur func(cur *ListNode) *ListNode
	checkRecur = func(cur *ListNode) *ListNode {
		if cur != nil {
			cur.Next = checkRecur(cur.Next)
			num++
			if num == n {
				return cur.Next
			}
		}
		return cur
	}

	return checkRecur(head)
}

// 双指针
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	slow, fast := sentinel, sentinel
	for i := 0; i < n+1; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return sentinel.Next
}
