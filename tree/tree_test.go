package tree

import (
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	bt := CreateBTree()
	bt.Insert(nil, 1, 0)
	bt.Insert(1, 2, 0)
	bt.Insert(1, 3, 1)
	bt.Insert(3, 4, 0)

	arr := inorderTraversal(bt.root)

	t.Log(arr)
}

func Test_isSameTree(t *testing.T) {
	t1 := CreateBTree()
	t1.Insert(nil, 1, 0)
	t1.Insert(1, 2, 0)

	t2 := CreateBTree()
	t2.Insert(nil, 1,  0)
	t2.Insert(1, 2,  1)

	t.Log(isSameTree1(t1.root, t2.root))
	t.Log(isSameTree2(t1.root, t2.root))
}

//func TestBTree_LevelSearch(t *testing.T) {
//	bt := CreateBTree()
//	bt.InsertByLevelSearch(nil, "a", 0)
//	bt.InsertByLevelSearch("a", "b", 0)
//	bt.InsertByLevelSearch("a", "d", 1)
//	bt.InsertByLevelSearch("b", "c", 1)
//	bt.InsertByLevelSearch("d", "e", 0)
//	bt.InsertByLevelSearch("d", "f", 1)
//
//	bt.InOrderTraverse(bt.root)
//	println()
//	bt.PreOrderTraverse(bt.root)
//	println()
//	bt.PostOrderTraverse(bt.root)
//	println()
//	bt.LevelTraverse()
//	println()
//}
//
//func TestRebuildBinaryTree(t *testing.T) {
//	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
//	inorder  := []int{4, 7, 2, 1, 5, 3, 8, 6}
//
//	tree := RebuildBinaryTree(preorder, inorder)
//
//	pre := printPreOrder(tree)
//	assert.Equal(t, preorder, pre)
//	t.Log(pre)
//
//	in := printInOrder(tree)
//	assert.Equal(t, inorder, in)
//	t.Log(in)
//
//	preorder = []int{1, 2, 3, 4}
//	inorder  = []int{3, 4, 2, 1}
//	tree = RebuildBinaryTree(preorder, inorder)
//	assert.Equal(t, preorder, printPreOrder(tree))
//	assert.Equal(t, inorder, printInOrder(tree))
//}
//
//func Test_inorderTraversal(t *testing.T)  {
//	inorderTraversal()
//}
