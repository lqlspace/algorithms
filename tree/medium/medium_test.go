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
