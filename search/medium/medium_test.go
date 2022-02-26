package medium

import (
	"testing"
)

func Test_search(t *testing.T)  {
	t.Log(search([]int{4,5,6,7,0,1,2}, 0))
	t.Log(search([]int{4,5,6,7,0,1,2}, 3))
	t.Log(search([]int{1}, 0))
	t.Log(search([]int{3, 1}, 1))
}
