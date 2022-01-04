package queue

import (
	"container/list"
)

type CQueue struct {
	stackIn *list.List
	stackOut *list.List
}

func ConstructorCQueue() *CQueue {
	return &CQueue{
		stackIn:  list.New(),
		stackOut: list.New(),
	}
}

func (c *CQueue) AppendTail(val int) {
	c.stackIn.PushBack(val)
}

func (c *CQueue) DeleteHead() int {
	if c.stackOut.Len() == 0 {
		for c.stackIn.Len() > 0 {
			c.stackOut.PushBack(c.stackIn.Remove(c.stackIn.Back()))
		}
	}

	if c.stackOut.Len() == 0 {
		return -1
	}

	e := c.stackOut.Back()
	c.stackOut.Remove(e)

	return e.Value.(int)
}
