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
