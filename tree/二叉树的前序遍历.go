package tree

func preorderTraversal(root *TreeNode) (arr []int) {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		arr = append(arr, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)

	return
}

