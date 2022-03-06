package simple


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
	front, back := head, head
	for i := 0; i < k; i++ {
		if front == nil {
			return nil
		}
		front = front.Next
	}

	for front != nil {
		front = front.Next
		back = back.Next
	}

	return back
}
