package hard

import (
	"testing"
)

func Test_maxSlidingWindow(t *testing.T)  {
	t.Log(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	t.Log(maxSlidingWindow2([]int{1,3,-1,-3,5,3,6,7}, 3))
	t.Log(maxSlidingWindow([]int{1}, 1))
	t.Log(maxSlidingWindow2([]int{1}, 1))
	t.Log(maxSlidingWindow([]int{5, 3, 4}, 1))
}
