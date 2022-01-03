package heap

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pqueue := NewPriorityQueue(4)
	pqueue.Push(Node{priority:3, value:3})
	pqueue.Push(Node{priority:4, value:4})
	pqueue.Push(Node{priority:1, value:1})
	pqueue.Push(Node{priority:9, value:9})

	node := pqueue.Pop()
	fmt.Printf("val = %d\n", node.value)

	node = pqueue.Pop()
	fmt.Printf("val = %d\n", node.value)

	pqueue.Push(Node{priority:5, value: 5})
	node = pqueue.Pop()
	fmt.Printf("val = %d\n", node.value)
}
