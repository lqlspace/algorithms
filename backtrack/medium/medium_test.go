package medium

import (
	"testing"
)

func Test_permute(t *testing.T)  {
	t.Log(permute([]int{1, 2, 3}))
}

func Test_restoreIpAddresses(t *testing.T)  {
	t.Log(restoreIpAddresses("1001125"))
	t.Log(restoreIpAddresses("25525511135"))
	t.Log(restoreIpAddresses("0000"))
	t.Log(restoreIpAddresses("101023"))
}

func Test_combinationSum(t *testing.T) {
	t.Log(combinationSum([]int{2,3,6,7}, 7))
	t.Log(combinationSum([]int{2,3,5}, 8))
	t.Log(combinationSum([]int{2}, 1))
}
