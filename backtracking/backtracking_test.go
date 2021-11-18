package backtracking

import (
	"testing"
)

// 八皇后问题
func TestCal8Queue(t *testing.T) {
	Cal8Queue(0)
}


// 0-1背包问题
func TestBag(t *testing.T) {
	items := []int{3, 5, 13, 15, 25}
	var maxWeight = -1
	Bag(0, 0, 21, items, &maxWeight)

	println(maxWeight)
}


