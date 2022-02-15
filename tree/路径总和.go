package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(root, 0, targetSum)
}

// 深度优先
// 时间复杂度O(N), 空间复杂度O(N)
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

//  广度优先搜索
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queNode := []*TreeNode{root}
	queVal := []int{root.Val}

	for len(queNode) > 0 {
		node := queNode[0]
		queNode = queNode[1:]
		val := queVal[0]
		queVal = queVal[1:]

		if node.Left == nil && node.Right == nil {
			if val == targetSum {
				return true
			}
			continue
		}

		if node.Left != nil {
			queNode = append(queNode, node.Left)
			queVal = append(queVal, val + node.Left.Val)
		}

		if node.Right != nil {
			queNode = append(queNode,  node.Right)
			queVal = append(queVal, val + node.Right.Val)
		}
	}

	return false
}
