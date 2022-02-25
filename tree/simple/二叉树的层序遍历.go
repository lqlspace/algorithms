package simple

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sz := len(queue)
		var arr  []int
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans = append(ans, arr)
	}

	return ans
}
