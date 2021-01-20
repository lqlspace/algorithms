package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 数组法（时间复杂度O(N),空间复杂度O(N)）
func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return true
	}
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
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
	if head == nil || head.Next == nil {
		return true
	}

	faster, slower := head, head
	for faster != nil && faster.Next != nil {
		slower = slower.Next
		faster = faster.Next.Next
	}

	rev := reverseLink(slower)

	for rev != nil || rev != slower {
		if head.Val != rev.Val {
			return false
		}
		rev = rev.Next
	}

	return true
}

func reverseLink(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	sentinel := new(ListNode)
	for node != nil {
		p := node
		node = node.Next
		p.Next = sentinel.Next
		sentinel.Next = p
	}

	return sentinel.Next
}
