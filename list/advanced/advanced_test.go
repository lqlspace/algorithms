package advanced

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T)  {
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 2)))
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 3)))
	t.Log(TraverseList(reverseKGroup(CreateList([]int{1,2,3,4,5}), 1)))
}
