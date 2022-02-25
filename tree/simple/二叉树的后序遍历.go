package simple

func postorderTraversal(root *TreeNode) (arr []int) {
	var post func(root *TreeNode)
	post = func(root *TreeNode) {
		if root == nil {
			return
		}

		post(root.Left)
		post(root.Right)
		arr = append(arr, root.Val)
	}

	post(root)

	return
}
