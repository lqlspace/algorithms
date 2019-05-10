package sorting

import (
	"testing"
	"algorithm-data-structure/algorithms/rand"
)


func TestHeapSort(t *testing.T) {
	items := randnum.RandomIntSlice(30)
	t.Log(items)
	HeapSort(items)
	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			t.Error("[error] heap sort error!")
			return
		}
	}
	t.Log(items)
	t.Log("heap sort succeed!")
}
