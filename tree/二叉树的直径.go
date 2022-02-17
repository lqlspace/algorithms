package tree

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
