package stack

import (
	"fmt"
)

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (as *ArrayStack) IsEmpty() bool {
	if as.top < 0 {
		return  true
	}

	return false
}

func (as *ArrayStack) Push(data interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, data)
	} else {
		as.data[as.top] = data
	}
}

func (as *ArrayStack) Pop() interface{} {
	if  as.IsEmpty() {
		return nil
	}

	data := as.data[as.top]
	as.top--
	return data
}

func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}

	return as.data[as.top]
}

func (as *ArrayStack) Resume() {
	as.top = -1
}

func (as *ArrayStack) Print() {
	if as.IsEmpty() {
		fmt.Println("empty stack")
		return
	}

	for i := as.top; i >= 0; i-- {
		fmt.Println(as.data[i])
	}
}
