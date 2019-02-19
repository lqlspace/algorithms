package sorting

import (
	"testing"
	"algorithm-data-structure/algorithms/rand"
)

func TestBitmapSort(t *testing.T) {
	items := randnum.RandomDistinctIntSlice(40)
//	items := []int{3, 2, 9, 10, 13, 14, 17, 11, 4, 5}
	t.Log(items)
	BitmapSort(items, 0, 100)
	t.Log(items)
	for i := 0; i < len(items) - 1; i++ {
		if items[i] > items[i+1] {
			t.Error("[error] bitmap sort error!")
			return
		}
	}

	t.Log("bitmap sort succeed!")
}
