package medium


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel1 := &ListNode{Next:head}
	cur := sentinel1
	sentinel2 := new(ListNode)
	var last *ListNode
	var num int
	for cur.Next != nil {
		num++
		if m <= num && num <= n {
			node := cur.Next
			cur.Next = cur.Next.Next
			node.Next = sentinel2.Next
			sentinel2.Next = node
			if last == nil {
				last = node
			}
		} else if num > n {
			last.Next = cur.Next
			cur.Next = sentinel2.Next
			break
		} else {
			cur = cur.Next
		}
	}
	if cur.Next == nil {
		last.Next = cur.Next
		cur.Next = sentinel2.Next
	}

	return sentinel1.Next
}
