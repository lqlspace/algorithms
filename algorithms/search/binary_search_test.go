package search

import (
	"testing"
)


func TestBinarySearch(t *testing.T) {
	si := NewSliceInt(1, 2, 3, 6, 7)
	idx, e := si.BinarySearch(7)
	if e {
		t.Logf("exist, location:%v",idx)
	} else {
		t.Logf("not exist!")
	}
}
