package simple

func findMode(root *TreeNode) (ans []int) {
	ans = []int{}
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}

		if count == maxCount {
			ans = append(ans, base)
		} else if count > maxCount {
			maxCount = count
			ans = []int{base}
		}
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return
}
