package medium

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	leftLen := len(inorder[:i])
	root.Left = buildTree(preorder[1:leftLen+1], inorder[:i])
	root.Right = buildTree(preorder[leftLen+1:], inorder[i+1:])

	return root
}
