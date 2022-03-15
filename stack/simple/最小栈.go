package simple

import (
	"math"
)

type Item struct {
	Val int
	Min int
}

type MinStack struct {
	items []*Item
}


func Constructor() MinStack {
	return MinStack{
		items: []*Item{},
	}
}


func (this *MinStack) Push(val int)  {
	item := &Item{Val: val, Min: val}
	if len(this.items) > 0 {
		item.Min = min(this.GetMin(), val)
	}
	this.items = append(this.items, item)
}


func (this *MinStack) Pop()  {
	if len(this.items) == 0 {
		return
	}
	this.items = this.items[:len(this.items)-1]
}


func (this *MinStack) Top() int {
	if len(this.items) > 0 {
		return this.items[len(this.items)-1].Val
	}

	return math.MinInt64
}


func (this *MinStack) GetMin() int {
	if len(this.items) > 0 {
		return this.items[len(this.items)-1].Min
	}

	return math.MinInt64
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
