package sort

import (
	"testing"
	"algorithm-data-structure/algorithms/rand"
)


func TestBubbleSort(t *testing.T) {
	items := randnum.RandomIntSlice(20)
	BubbleSort(items)
	for i := 0; i < 19; i++ {
		if items[i] > items[i+1] {
			t.Error("[error] sort fail!")
			return
		}
	}
}
