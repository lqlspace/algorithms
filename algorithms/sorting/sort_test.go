package sort

import (
	"testing"
)

func TestInsertionSortIntSlice(t *testing.T) {
	arr := IntSlice{3, 3, 9, 6, 2, 7, 7, 8, 1, 0}
	arr.Sort()
	t.Log(arr)
}
