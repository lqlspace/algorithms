package medium

func sumNumbers(root *TreeNode) int {
	return dfsPreSum(root, 0)
}

func dfsPreSum(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}

	preSum = preSum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return preSum
	}

	return dfsPreSum(root.Left, preSum) + dfsPreSum(root.Right, preSum)
}
