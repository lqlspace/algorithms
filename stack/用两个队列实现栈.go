package stack

import (
	"container/list"
)

type MyStack struct {
	queueIn *list.List
	queueOut *list.List
}

func ConstructorMyStack() MyStack {
	return MyStack{
		queueIn: list.New(),
		queueOut: list.New(),
	}
}

func (s *MyStack) Push(val int) {
	s.queueIn.PushBack(val)
	for s.queueOut.Len() > 0 {
		s.queueIn.PushBack(s.queueOut.Remove(s.queueOut.Front()))
	}
	s.queueIn, s.queueOut = s.queueOut, s.queueIn
}

func (s *MyStack) Pop() int {
	if s.queueOut.Len() == 0 {
		return -1
	}

	return s.queueOut.Remove(s.queueOut.Front()).(int)
}

func (s *MyStack) Top() int {
	if s.queueOut.Len() == 0 {
		return -1
	}

	return s.queueOut.Front().Value.(int)
}

func (s *MyStack) Empty() bool {
	return s.queueOut.Len() == 0
}
