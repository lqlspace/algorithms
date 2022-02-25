package medium

import (
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T)  {
	head := CreateBTree([]int{3,9,20,null,null,15,7})
	t.Log(zigzagLevelOrder(head))
}
