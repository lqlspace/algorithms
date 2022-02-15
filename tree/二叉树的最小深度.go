package tree


// 广度优先搜索
// 时间复杂度O(N), 空间复杂度O(N)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	ans := 1
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return ans
			}
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

// 深度优先搜索
// 时间复杂度O(N), 空间复杂度O(H)
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return min(minDepth2(root.Left), minDepth2(root.Right)) + 1
}

func min(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a < b {
		return a
	}

	return b
}
