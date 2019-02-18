package sorting

import (
	"testing"
	randnum "algorithm-data-structure/algorithms/rand"
	"fmt"
)

func TestPartition(t *testing.T) {
	items := randnum.RandomIntSlice(20)
	fmt.Println(items)
	pivotIdx := Partition(items, 0, len(items)-1)
	fmt.Println(items, pivotIdx)

	for i := 0; i < len(items); i++ {
		if (i < pivotIdx && items[i] > items[pivotIdx]) || (i > pivotIdx && items[i] < items[pivotIdx]) {
			t.Error("[error] partition error!")
			return 
		}
	}

}

func TestQuickSort(t *testing.T) {
	items := randnum.RandomIntSlice(30)
	t.Log(items)
	QuickSort(items, 0, len(items)-1)
	t.Log(items)
}
