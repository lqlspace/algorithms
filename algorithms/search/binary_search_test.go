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


func TestLineSearch(t *testing.T) {
	si := NewSliceInt(1, 2, 3, 6, 7)
	idx, e := si.LineSearch(7)
	if e {
		t.Logf("exist, location: %v", idx)
	} else {
		t.Logf("not exist!")
	}
}


func TestBitmapSearch(t *testing.T) {
	si := NewSliceIntBitmap(1, 2, 3, 6, 7)
	t.Log(si)
	e := si.BitmapSearch(7)
	if e {
		t.Logf("exist!")
	} else {
		t.Logf("not exist!")
	}
	e = si.BitmapSearch(3)
	if e {
		t.Logf("exist!")
	} else {
		t.Logf("not exist!")
	}
}
