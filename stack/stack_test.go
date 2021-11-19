package stack

import (
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	stack := NewArrayStack()

	t.Log(stack.IsEmpty())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	t.Log(stack.Top())

	stack.Print()

	stack.Pop()
	stack.Pop()

	println()

	stack.Print()
}
