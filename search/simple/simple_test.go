package simple

import (
	"testing"
)

func Test_search(t *testing.T)  {
	t.Log(search([]int{-1,0,3,5,9,12}, 9))
	t.Log(search2([]int{-1,0,3,5,9,12}, 9))
	t.Log(search([]int{-1,0,3,5,9,12}, 2))
	t.Log(search2([]int{-1,0,3,5,9,12}, 2))
}

func Test_mySqrt(t *testing.T)  {
	t.Log(mySqrt(4))
	t.Log(mySqrt2(4))
	t.Log(mySqrt(8))
	t.Log(mySqrt2(8))
	t.Log(mySqrt(9))
	t.Log(mySqrt2(9))
	t.Log(mySqrt(15))
	t.Log(mySqrt2(15))
	t.Log(mySqrt(16))
	t.Log(mySqrt2(16))
}
