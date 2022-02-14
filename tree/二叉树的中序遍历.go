package tree

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var arr []int
	left := inorderTraversal(root.Left)
	arr = append(arr, left...)
	arr = append(arr, root.Val)
	right := inorderTraversal(root.Right)
	arr = append(arr, right...)

	return arr
}
