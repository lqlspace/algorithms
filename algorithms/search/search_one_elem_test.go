package search

import (
	"testing"
)


func TestBinarySearch(t *testing.T) {
	testData := GetRandNums("sorted.txt")
	
	si := NewSliceInt(testData...)
	idx, e := si.BinarySearch(6700)
	if e {
		t.Logf("exist, location:%v",idx)
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, e)
	}
}


func TestLineSearch(t *testing.T) {
	testData := GetRandNums("sorted.txt")

	si := NewSliceInt(testData...)
	idx, e := si.LineSearch(6700)
	if e {
		t.Logf("exist, location: %v", idx)
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, e)
	}
}


func TestBitmapSearch(t *testing.T) {
	testData := GetRandNums("sorted.txt")

	si := NewSliceIntBitmap(testData...)

	e := si.BitmapSearch(6700)
	if e {
		t.Logf("exist!!!")
	} else {
		t.Errorf("expected: %t, actual: %t", true, e)
	}
}


func BenchmarkLineSearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNums("sorted.txt")
	si := NewSliceInt(testData...)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		si.LineSearch(6700)
	}
}


func BenchmarkBinarySearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNums("sorted.txt")
	si := NewSliceInt(testData...)
	b.StartTimer()


	for i := 0; i < b.N; i++ {
		si.BinarySearch(6700)
	}
}



func BenchmarkBitmapSearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNums("sorted.txt")
	si := NewSliceIntBitmap(testData...)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		si.BitmapSearch(6700)
	}
}
