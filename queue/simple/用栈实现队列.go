package simple

import (
	"container/list"
)

type Stack struct {
	items []int
}

func (s *Stack) Empty() bool {
	return s.Size() ==  0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	val := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]

	return val
}

func (s *Stack) Top() int {
	return s.items[s.Size()-1]
}


type MyQueue struct {
	in, out Stack
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.in.Push(x)
}


func (this *MyQueue) Pop() int {
	if !this.out.Empty() {
		return this.out.Pop()
	}

	for !this.in.Empty() {
		v := this.in.Pop()
		this.out.Push(v)
	}

	return this.out.Pop()
}


func (this *MyQueue) Peek() int {
	if !this.out.Empty() {
		return this.out.Top()
	}

	for !this.in.Empty() {
		v := this.in.Pop()
		this.out.Push(v)
	}

	return this.out.Top()
}


func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}


// 方法2：用标准库
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
