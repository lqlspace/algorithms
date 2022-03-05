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

		maxLeftSide := recur(root.Left)
		maxRightSide := recur(root.Right)
		rootPath := max(maxLeftSide, 0) + max(maxRightSide, 0) + root.Val
		maxPath = max(rootPath, maxPath)

		return max(max(maxLeftSide, maxRightSide), 0)+root.Val
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
