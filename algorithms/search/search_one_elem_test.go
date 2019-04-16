package search

import (
	"testing"
)


func TestLineSearch(t *testing.T) {
	testData := GetRandNumsJSON("sorted.txt")

	si := NewSliceInt(testData...)
	idx, e := si.LineSearch(99267)
	if e {
		t.Logf("exist, location: %v", idx)
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, e)
	}
}



func TestBinarySearch(t *testing.T) {
	testData := GetRandNumsJSON("sorted.txt")
	
	si := NewSliceInt(testData...)
	idx, e := si.BinarySearch(99267)
	if e {
		t.Logf("exist, location:%v",idx)
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, e)
	}
}






func BenchmarkLineSearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNumsJSON("sorted.txt")
	si := NewSliceInt(testData...)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		si.LineSearch(99267)
	}
}


func BenchmarkBinarySearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNumsJSON("sorted.txt")
	si := NewSliceInt(testData...)
	b.StartTimer()


	for i := 0; i < b.N; i++ {
		si.BinarySearch(99267)
	}
}





/********************************链表***********************************/


func TestLineSearchList(t *testing.T) {
	datas := GetRandNumsJSON("sorted.txt")
	l := NewListInt(datas...)
	exist := l.LineSearch(99267)
	if !exist {
		t.Errorf("Error: expected: true, actual: %t", exist)
	} else {
		t.Logf("Exist : %t", exist)
	}
}	


func BenchmarkLineSearchList(b *testing.B) {
	b.StopTimer()
	datas := GetRandNumsJSON("sorted.txt")
	l := NewListInt(datas...)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		l.LineSearch(99267)
	}
}


/*********************************bitmap****************************************/

func TestBitmapSearch(t *testing.T) {
	testData := GetRandNumsJSON("sorted.txt")

	si := NewSliceIntBitmap(testData...)

	e := si.BitmapSearch(99267)
	if e {
		t.Logf("exist!!!")
	} else {
		t.Errorf("expected: %t, actual: %t", true, e)
	}
}


func BenchmarkBitmapSearch(b *testing.B) {
	b.StopTimer()
	testData := GetRandNumsJSON("sorted.txt")
	si := NewSliceIntBitmap(testData...)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		si.BitmapSearch(99267)
	}
}
