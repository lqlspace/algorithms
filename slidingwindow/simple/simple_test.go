package simple

import (
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T)  {
	t.Log(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	t.Log(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	t.Log(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
}
