package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestConstructorMyStack(t *testing.T) {
	s := ConstructorMyStack()
	s.Push(1)
	v := s.Pop()
	assert.Equal(t, v, 1)

	s.Push(2)
	s.Push(3)
	s.Push(4)
	assert.Equal(t, s.Pop(), 4)
	assert.Equal(t, s.Pop(), 3)
	assert.Equal(t, s.Pop(), 2)

}
