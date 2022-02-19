package dp

import (
	"testing"
)

func Test_maxSubArray(t *testing.T)  {
	arr := []int{2,1,-3,4,-1,2,1,-5,4}
	t.Log(maxSubArray(arr))
}

func Test_climbStirs(t *testing.T) {
	t.Log(climbStairs(2))
	t.Log(climbStairs2(2))
	t.Log(climbStairs(3))
	t.Log(climbStairs2(3))
}

func Test_maxProfit(t *testing.T) {
	arr :=  []int{7,1,5,3,6,4}
	t.Log(maxProfit(arr))

	arr =  []int{7,6,4,3,1}
	t.Log(maxProfit(arr))
}

func Test_generate(t *testing.T) {
	t.Log(generate(1))
	t.Log(generate(2))
	t.Log(generate(3))
	t.Log(generate(4))
	t.Log(generate(5))
}
