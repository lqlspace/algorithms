package simple

import (
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T)  {
	t.Log(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	t.Log(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	t.Log(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
}

func Test_findMaxAverage(t *testing.T) {
	t.Log(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))
	t.Log(findMaxAverage([]int{5}, 1))
}
