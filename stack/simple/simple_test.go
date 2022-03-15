package simple

import (
	"testing"
)

func Test_isValid(t *testing.T)  {
	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("([)]"))
	t.Log(isValid("[()]"))
	t.Log(isValid("[()"))
	t.Log(isValid("(("))
}

func Test_minStack(t *testing.T) {
	minStack := Constructor()
	minStack.Pop()
	minStack.GetMin()
	minStack.Push(4)
	minStack.Push(3)
	minStack.Push(-1)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.GetMin())
}
