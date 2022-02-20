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

func Test_getRow(t *testing.T) {
	t.Log(getRow(0))
	t.Log(getRow(1))
	t.Log(getRow(2))
	t.Log(getRow(3))
	t.Log(getRow(4))
	t.Log(getRow(5))
}

func Test_countBits(t *testing.T)  {
	t.Log(countBits(7))
	t.Log(countBits2(7))
	t.Log(countBits3(7))
}

func Test_isSubsequence(t *testing.T)  {
	t.Log(isSubsequence("ace", "abcdfe"))
	t.Log(isSubsequence2("ace", "abcdfe"))
	t.Log(isSubsequence("aaabbb", "afsdfadfsadfbbdsfsdbfsdb"))
	t.Log(isSubsequence2("aaabbb", "afsdfadfsadfbbdsfsdbfsdb"))
	t.Log(isSubsequence("aaabbbc", "afsdfadfsadfbbdsfsdbfsdb"))
	t.Log(isSubsequence2("aaabbbc", "afsdfadfsadfbbdsfsdbfsdb"))
}

func Test_fib(t *testing.T) {
	t.Log(fib(2))
	t.Log(fib2(2))
	t.Log(fib(3))
	t.Log(fib2(3))
	t.Log(fib(4))
	t.Log(fib2(4))
	t.Log(fib(5))
	t.Log(fib2(5))
	t.Log(fib(6))
	t.Log(fib2(6))
	t.Log(fib(7))
	t.Log(fib2(7))
}

func Test_minCostClimbingStairs(t *testing.T) {
	t.Log(minCostClimbingStairs([]int{10, 15, 20}))
	t.Log(minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1}))
}
