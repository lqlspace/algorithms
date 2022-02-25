package simple

// 深度优先搜索
func sumOfLeftLeaves(root *TreeNode) int {
	if root ==  nil {
		return 0
	}

	return dfsSum(root)
}

func dfsSum(root *TreeNode) int {
	var ans int

	if root.Left != nil {
		if isLeaveNode(root.Left) {
			ans += root.Left.Val
		} else {
			ans += dfsSum(root.Left)
		}
	}

	if root.Right != nil && !isLeaveNode(root.Right) {
		ans += dfsSum(root.Right)
	}

	return  ans
}

func isLeaveNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}


func sumOfLeftLeaves2(root *TreeNode) int {
	var ans int
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if isLeaveNode(node.Left) {
				ans += node.Left.Val
			} else {
				queue = append(queue, node.Left)
			}
		}

		if node.Right != nil && !isLeaveNode(node.Right) {
			queue = append(queue, node.Right)
		}
	}

	return ans
}
