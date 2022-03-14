package medium

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int

	var dfsPathSum func(root *TreeNode, remainder int)
	dfsPathSum = func(root *TreeNode, remainder int) {
		if root == nil {
			return
		}

		remainder -= root.Val
		path = append(path, root.Val)
		defer func() {path = path[:len(path)-1]}()
		if root.Left == nil && root.Right == nil && remainder == 0 {
			ans = append(ans, append([]int(nil), path...))
		}

		dfsPathSum(root.Left, remainder)
		dfsPathSum(root.Right, remainder)
	}

	dfsPathSum(root, targetSum)

	return ans
}
