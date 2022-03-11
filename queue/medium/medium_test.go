package medium

import (
	"testing"
)

func Test_myCircularQueue(t *testing.T)  {
	queue := Constructor(3)
	t.Log(queue.IsEmpty())
	t.Log(queue.EnQueue(1))
	t.Log(queue.EnQueue(2))
	t.Log(queue.EnQueue(3))
	t.Log(queue.EnQueue(4))
	t.Log(queue.IsFull())
	t.Log(queue.IsEmpty())
	t.Log(queue.Front())
	t.Log(queue.Rear())
	t.Log(queue.DeQueue())
	t.Log(queue.IsFull())
	t.Log(queue.Front())
	t.Log(queue.Rear())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.IsEmpty())
	t.Log(queue.Front())
	t.Log(queue.Rear())
	t.Log(queue.EnQueue(4))
	t.Log(queue.Front())
	t.Log(queue.Rear())

}
