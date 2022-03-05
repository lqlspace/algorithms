package hard

import (
	"math"
)

func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt64
	var recur func(root *TreeNode) int
	recur = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		maxLeft := max(recur(root.Left), 0)
		maxRight := max(recur(root.Right), 0)

		rootPath := maxLeft + maxRight + root.Val
		maxPath = max(rootPath, maxPath)

		return max(maxLeft, maxRight)+root.Val
	}

	recur(root)

	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
