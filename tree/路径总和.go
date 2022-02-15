package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(root, 0, targetSum)
}

func pathSum(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val
	if root.Left == nil && root.Right == nil {
		return  sum  == targetSum
	}

	return pathSum(root.Left, sum, targetSum) || pathSum(root.Right, sum, targetSum)
}
