package other

import (
	"testing"
)

func TestMaxProfit(t *testing.T)  {
	profits := []int{3, 2, 1, 7, 6}
	v :=  maxProfit2(profits)
	t.Log(v)
}

func TestCanJump(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	v := canJump(nums)
	t.Log(v)

	nums = []int{2,3,1,1,4}
	v  = canJump(nums)
	t.Log(v)
}
