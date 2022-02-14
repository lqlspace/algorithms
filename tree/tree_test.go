package tree

import (
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	node := CreateBTree([]int{1, 2, 3, null, null, 4})

	arr := inorderTraversal(node)

	t.Log(arr)
}

func Test_isSameTree(t *testing.T) {
	node1 := CreateBTree([]int{1, 2})
	node2 := CreateBTree([]int{1, null, 2})

	t.Log(isSameTree1(node1, node2))
	t.Log(isSameTree2(node1, node2))
}

func Test_isSymmetric(t *testing.T) {
	bt := CreateBTree([]int{1, 2, 2, 2, null, 2})
	t.Log(isSymmetric(bt))

	bt = CreateBTree([]int{1, 2, 2, 3, 4, 4, 3})
	t.Log(isSymmetric(bt))

	bt = CreateBTree([]int{5, 4, 1, null, 1, null, 4, 2, null, 2, null})
	t.Log(isSymmetric(bt))
}

func Test_maxDepth(t *testing.T) {
	bt := CreateBTree([]int{3, 9, 20, null, null, 15, 7})
	t.Log(maxDepth(bt))
	t.Log(maxDepth2(bt))
}
