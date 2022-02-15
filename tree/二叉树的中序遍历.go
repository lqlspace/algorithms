package tree

func inorderTraversal(root *TreeNode) (arr []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		arr =append(arr, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return
}
