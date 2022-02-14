package tree

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如：输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建如图所示的二叉树并输出它的头节点。
 */

//func RebuildBinaryTree(preorder, inorder []int) *BinaryTreeNode {
//	if len(preorder) == 0 || len(inorder) == 0 {
//		return nil
//	}
//
//	return rebuildTreeRecur(preorder, 0, len(preorder)-1, inorder, 0,  len(inorder)-1)
//}
//
//func rebuildTreeRecur(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *BinaryTreeNode {
//	if preStart > preEnd || inStart > inEnd {
//		return nil
//	}
//
//	root := new(BinaryTreeNode)
//	root.Value = preorder[preStart]
//
//	var rootIdx int
//	for i := range inorder {
//		if inorder[i] == root.Value {
//			rootIdx = i
//			break
//		}
//	}
//
//	leftCount := rootIdx - inStart
//	root.Left = rebuildTreeRecur(preorder, preStart+1, preStart+leftCount, inorder, inStart, rootIdx-1)
//	root.Right = rebuildTreeRecur(preorder, preStart+1+leftCount, preEnd, inorder, rootIdx+1, inEnd)
//
//	return  root
//}
//
//func printPreOrder(tree *BinaryTreeNode) []int {
//	if tree == nil {
//		return nil
//	}
//
//	var res []int
//	res = append(res, tree.Value)
//	res = append(res, printPreOrder(tree.Left)...)
//	res = append(res, printPreOrder(tree.Right)...)
//
//	return res
//}
//
//func printInOrder(tree *BinaryTreeNode) []int {
//	if tree == nil {
//		return  nil
//	}
//
//	var res []int
//	res = append(res, printInOrder(tree.Left)...)
//	res = append(res, tree.Value)
//	res = append(res, printInOrder(tree.Right)...)
//
//	return res
//}
