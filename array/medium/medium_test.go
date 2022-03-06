package medium

import (
	"testing"
)

func Test_threeSum(t *testing.T)  {
	t.Log(threeSum([]int{-1, 2, 0, 1}))
}

func Test_spiralOrders(t *testing.T) {
	t.Log(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	t.Log(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))

	t.Log(spiralOrder([][]int{
		{6, 9, 7},
	}))
}

func Test_merge(t *testing.T)  {
	t.Log(merge([][]int{
		{1, 3},{2, 6}, {8, 10}, {15, 18},
	}))
}
