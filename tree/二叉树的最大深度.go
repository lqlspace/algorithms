package tree


// 1. 深度优先遍历
// 时间复杂度O(N), 空间复杂度O(height)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}



// 2. 迭代法
// 时间复杂度O(N), 空间复杂度最坏情况O(N)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}

	return ans
}

