package hard

import (
	"testing"
)

func Test_maxPathSum(t *testing.T)  {
	head := CreateBTree([]int{1, 2, 3})
	t.Log(maxPathSum(head))

	head = CreateBTree([]int{-10,9,20,null,null,15,7})
	t.Log(maxPathSum(head))

	head = CreateBTree([]int{2, -1})
	t.Log(maxPathSum(head))

	head = CreateBTree([]int{9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6})
	t.Log(maxPathSum(head))
}
