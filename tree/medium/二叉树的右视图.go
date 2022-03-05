package medium

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sz, first := len(queue), true
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if first {
				ans = append(ans, node.Val)
				first = false
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			sz--
		}
	}

	return ans
}
