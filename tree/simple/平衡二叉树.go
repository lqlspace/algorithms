package simple

// 时间复杂度O(N^2), 空间复杂度O(N^2)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	n := left - right
	if n > 1 || n < -1 {
		return false
	}

	return true
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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
