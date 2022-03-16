package medium

import (
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T)  {
	head := CreateBTree([]int{3,9,20,null,null,15,7})
	t.Log(zigzagLevelOrder(head))
}

func Test_lowestCommonAncestor(t *testing.T)  {
	head := CreateBTree([]int{3,5,1,6,2,0,8,null,null,7,4})
	p, _ := dfs(head, 5)
	q, _ := dfs(head, 1)
	t.Log(lowestCommonAncestor(head, p, q))
}

func Test_rightSideView(t *testing.T)  {
	head := CreateBTree([]int{1,2,3,null,5,null,4})
	t.Log(rightSideView(head))

	head = CreateBTree([]int{1,null,3})
	t.Log(rightSideView(head))

	head = CreateBTree([]int{})
	t.Log(rightSideView(head))
}

func Test_buildTree(t *testing.T) {
	t.Log(LevelTraversal(buildTree([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})))
}

func Test_sumNumbers(t *testing.T) {
	t.Log(sumNumbers(CreateBTree([]int{1, 2, 3})))
	t.Log(sumNumbers(CreateBTree([]int{4,9,0,5,1})))
}

func Test_pathSum(t *testing.T) {
	head := CreateBTree([]int{5,4,8,11,null,13,4,7,2,null,null,5,1})
	t.Log(pathSum(head, 22))

	head = CreateBTree([]int{1, 2, 3})
	t.Log(pathSum(head, 5))

	head = CreateBTree([]int{1, 2})
	t.Log(pathSum(head, 0))
}

func Test_isValidBST(t *testing.T) {
	head := CreateBTree([]int{2, 1, 3})
	t.Log(isValidBST(head))

	head = CreateBTree([]int{5,1,4,null,null,3,6})
	t.Log(isValidBST(head))

	head = CreateBTree([]int{0})
	t.Log(isValidBST(head))
}
