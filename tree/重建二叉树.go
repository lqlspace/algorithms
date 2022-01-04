package tree

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如：输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建如图所示的二叉树并输出它的头节点。
 */


func RebuildBinaryTree(preorder, inorder []int) *BinaryTreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	return createTreeRecur(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func createTreeRecur(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *BinaryTreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	root := &BinaryTreeNode{Value:preorder[preStart]}
	var rootIdx int
	for i := range inorder {
		if inorder[i] == root.Value {
			rootIdx = i
			break
		}
	}

	leftCount := rootIdx - inStart
	root.Left = createTreeRecur(preorder, preStart+1, preStart+leftCount, inorder, inStart, rootIdx-1)
	root.Right = createTreeRecur(preorder, preStart+leftCount+1, preEnd, inorder, rootIdx+1, inEnd)

	return root
}


// 前序遍历
func printPreOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}
	order := append([]int(nil), root.Value)
	if root.Left != nil {
		order = append(order, printPreOrder(root.Left)...)
	}
	if root.Right != nil {
		order = append(order, printPreOrder(root.Right)...)
	}
	return order
}

// 中序遍历
func printInOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		res = append(res, printInOrder(root.Left)...)
	}
	res = append(res, root.Value)
	if root.Right != nil {
		res = append(res, printInOrder(root.Right)...)
	}
	return res
}
