package simple

// 自顶向下：时间复杂度O(N^2)，空间复杂度O(N^2)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(maxDep(root.Left)-maxDep(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDep(root.Left), maxDep(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

// 自底向上的递归
func isBalanced2(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
