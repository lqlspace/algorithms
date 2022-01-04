package queue

import (
	"container/list"
)

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
 */

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
