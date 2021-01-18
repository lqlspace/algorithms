package linkedlist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两次遍历（时间复杂度O(n), 空间复杂度O(1)）
func getKthFromEndMethod1(head *ListNode, k int) *ListNode {
	num, cur := 0, head
	for cur != nil {
		num++
		cur = cur.Next
	}
	if num < k {
		return nil
	}

	pos, cur := num - k, head
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	return cur
}

//双指针
func getKthFromEndMethod2(head *ListNode, k int) *ListNode {
	former, latter := head, head
	for i := 0; i < k; i++ {
		if former == nil {
			return nil
		}
		former = former.Next
	}

	for former != nil {
		former = former.Next
		latter = latter.Next
	}

	return latter
}
