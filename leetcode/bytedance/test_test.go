package bytedance

import (
	"fmt"
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

func TestSendRedPackage(t *testing.T) {
	arr := sendRedPackage(100, 5)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%.2f ", float64(arr[i])/100)
	}
}
