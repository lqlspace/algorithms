package heap

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{0, 3, 2, 1, 4, 5}
	Sort(arr, 5)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
