package simple

import (
	"testing"
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

