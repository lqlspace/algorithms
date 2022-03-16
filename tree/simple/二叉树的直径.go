package simple

import (
	"math"
)

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		ans = max(ans,  left+right+1)

		return max(left, right) + 1
	}

	depth(root)

	return ans - 1
}

func diameterOfBinaryTree2(root *TreeNode) int {
	ans := math.MinInt64

	var heightOfBTree func(root *TreeNode) int
	heightOfBTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := heightOfBTree(root.Left)
		rightHeight := heightOfBTree(root.Right)

		ans = max(ans, leftHeight + rightHeight)

		return max(leftHeight, rightHeight) + 1
	}

	heightOfBTree(root)

	return ans
}


