package medium

import (
	"math"
)

// 中序遍历（递归）
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64

	var inorderBST func(root *TreeNode) bool
	inorderBST = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := inorderBST(root.Left)
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		right := inorderBST(root.Right)

		return left && right
	}

	return inorderBST(root)
}
