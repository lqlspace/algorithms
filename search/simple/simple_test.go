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
