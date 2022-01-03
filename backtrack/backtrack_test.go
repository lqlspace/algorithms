package backtrack

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

// candidates = [2,3,5], target = 8
func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8


	res := CombinationSum(candidates, target)
	for _, item := range res {
		t.Log(item)
	}

	//candidates = [2,3,6,7], target = 7
	cs := []int{2, 3, 6, 7}
	target = 7
	res = CombinationSum(cs, target)
	for _, item := range res {
		t.Log(item)
	}
}

func TestCombinationSumII(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8

	res := CombinationSumII(candidates,  target)
	for _, item := range res {
		t.Log(item)
	}

	t.Log("******")

	candidates = []int{10,1,2,7,6,5}
	target = 8

	res = CombinationSumII(candidates,  target)
	for _, item := range res {
		t.Log(item)
	}

	t.Log("******")

	candidates = []int{2,5,1}
	target = 8

	res = CombinationSumII(candidates,  target)
	for _, item := range res {
		t.Log(item)
	}
}

