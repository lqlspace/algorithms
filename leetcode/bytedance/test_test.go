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

func TestMerge(t *testing.T) {
	a := []int{1,2,3,0,0,0}
	b := []int{2,5,6}
	merge(a, 3, b, 3)

	for _, v := range a {
		fmt.Printf("%d ", v)
	}

	a = []int{0}
	b = []int{1}
	merge(a, 0, b, 1)

	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}
