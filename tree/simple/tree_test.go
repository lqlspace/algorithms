package simple

import (
	"testing"
)

func Test_levelOrder(t *testing.T)  {
	head := CreateBTree([]int{3,9,20,null,null,15,7})
	t.Log(levelOrder(head))

	head = CreateBTree([]int{1})
	t.Log(levelOrder(head))

	head = CreateBTree([]int{})
	t.Log(levelOrder(head))
}
