package tree



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
