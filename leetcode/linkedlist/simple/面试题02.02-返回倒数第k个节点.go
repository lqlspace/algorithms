package simple

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法1：两次遍历（时间复杂度O(N)：常量因子为2,空间复杂度O(1)）
func kthToLast1(head *ListNode, k int) int {
	var num int
	for head != nil {
		num++
		head = head.Next
	}
	pos := num -k
	if pos < 0 {
		return 0
	}

	for i := 0; i < pos; i++ {
		head = head.Next
	}

	return head.Val
}

func kthToLast2(head *ListNode, k int) int {
	former, latter := head, head
	for i := 0; i < k; i++ {
		if former == nil {
			return 0
		}
		former = former.Next
	}

	for former != nil {
		former = former.Next
		latter = latter.Next
	}

	return latter.Val
}
