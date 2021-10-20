package other

import (
	"testing"
)

func TestMaxProfit(t *testing.T)  {
	profits := []int{3, 2, 1, 7, 6}
	v :=  maxProfit2(profits)
	t.Log(v)
}
