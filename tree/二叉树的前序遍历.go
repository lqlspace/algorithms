package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := []int{root.Val}
	if root.Left != nil {
		left := preorderTraversal(root.Left)
		arr = append(arr, left...)
	}
	if root.Right != nil {
		right := preorderTraversal(root.Right)
		arr = append(arr, right...)
	}

	return arr
}
