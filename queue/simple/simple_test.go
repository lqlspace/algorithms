package simple

import (
	"testing"
)

func Test_MyQueue(t *testing.T)  {
	q := Constructor()
	t.Log(q.Empty())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Empty())
	t.Log(q.Peek())
	t.Log(q.Pop())
	t.Log(q.Pop())
	q.Push(4)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Empty())
}
