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

func Test_searchRange(t *testing.T) {
	t.Log(searchRange([]int{1, 2, 2, 3}, 2))
	t.Log(searchRange([]int{1, 2, 3}, 2))
	t.Log(searchRange([]int{5,7,7,8,8,10}, 8))
	t.Log(searchRange([]int{5,7,7,8,8,10}, 6))
	t.Log(searchRange([]int{}, 0))
}
