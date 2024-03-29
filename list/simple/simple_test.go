package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateList(t *testing.T) {
	head := CreateList([]int{1, 2, 3, 4, 5})
	TraverseList(head)
}

func TestSwapNode(t *testing.T) {
	head := CreateList([]int{1, 2, 3, 4})
	head = swapPairs2(head)
	TraverseList(head)

	head = CreateList([]int{1, 2, 3})
	head = swapPairs2(head)
	TraverseList(head)

	head = CreateList([]int{1, 2})
	head = swapPairs2(head)
	TraverseList(head)

	head = CreateList([]int{1})
	head = swapPairs2(head)
	TraverseList(head)

	head = CreateList([]int{})
	head = swapPairs2(head)
	TraverseList(head)

	head = CreateList(nil)
	head = swapPairs2(head)
	TraverseList(head)
}

func TestRotateRight2(t *testing.T)  {
	head := CreateList([]int{1, 2, 3, 4})
	head = RotateRight2(head, 1)
	TraverseList(head)

	head = CreateList([]int{1, 2, 3, 4})
	head = RotateRight2(head, 3)
	TraverseList(head)

	head = CreateList([]int{1, 2, 3, 4})
	head = RotateRight2(head, 4)
	TraverseList(head)
}

func TestRotateRight3(t *testing.T)  {
	head := CreateList([]int{1, 2, 3, 4})
	head = RotateRight3(head, 1)
	TraverseList(head)

	head = CreateList([]int{1, 2, 3, 4})
	head = RotateRight3(head, 3)
	TraverseList(head)

	head = CreateList([]int{1, 2, 3, 4})
	head = RotateRight3(head, 4)
	TraverseList(head)
}

func TestPartition(t *testing.T) {
	head := CreateList([]int{3, 4, 9,  11, 2, 5})
	head = Partition(head, 1)
	TraverseList(head)
}

func Test_getKthFromEnd(t *testing.T) {
	node := getKthFromEnd(CreateList([]int{1, 2, 3, 4, 5}), 3)
	t.Log(node.Val)

	node = getKthFromEnd(CreateList([]int{1}), 3)
	assert.Nil(t, node)
}

