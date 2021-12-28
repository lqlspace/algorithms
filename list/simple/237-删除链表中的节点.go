package simple

// 此方法是在无法获取到当前节点的pre节点时采用的，且无法删除最后一个节点。
// （时间复杂度O(N), 空间复杂度O(N)）
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
