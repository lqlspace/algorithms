package sort

import (
	"testing"
)

func TestInsertionSortIntSlice(t *testing.T) {
	datas := IntSlice{3, -1, 3, 2, 5, 2, 7, 9}
	InsertionSort(datas, 0, len(datas))
	t.Log(datas)
}


func TestSelectionSortIntSlice(t *testing.T) {
	datas := IntSlice{-1, -3, 9, 9, 7, 6, 3, 5, 7, 8}
	SelectionSort(datas, 0, len(datas))
	t.Log(datas)
}


func TestShellSortIntSlice(t *testing.T) {
	datas := IntSlice{-1, 3, 2, 1, 2, 7, 6, 7, 9}
	ShellSort(datas, 0, len(datas))
	t.Log(datas)
}


func TestMergeSortIntSlice(t *testing.T) {
	datas := IntSlice{-3, 2, 9, 2, 1, 4, 6, 4, 5}
	MergeSort(datas)
	t.Log(datas)
}
