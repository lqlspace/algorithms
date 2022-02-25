package simple

func findTilt(root *TreeNode) int {
	var ans int

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		ans += abs(left - right)
		return left + right + node.Val
	}

	dfs(root)

	return ans
}
