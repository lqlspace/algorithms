package simple

/*
将小于X的数排在左边；
 */

func Partition(head *ListNode, x int) *ListNode {
	sentinel1, sentinel2 := new(ListNode), new(ListNode)
	sentinel1.Next = head
	cur, pre := sentinel1, sentinel2

	for cur.Next != nil {
		if cur.Next.Val > x {
			cur = cur.Next
			continue
		}

		node := cur.Next
		cur.Next = cur.Next.Next
		pre.Next = node
		pre = pre.Next
	}

	pre.Next = sentinel1.Next

	return sentinel2.Next
}
