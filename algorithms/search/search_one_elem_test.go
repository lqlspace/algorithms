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


/*****************************AVL Tree*********************************/

func TestDisplayAVLTree(t *testing.T) {
	datas := GetRandNumsJSON("sorted.txt")
	avl := NewAVLTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		avl = avl_insert(avl, datas[i]) 
	}

//	datasAsc := displayAsc(avl)
//	t.Log(datasAsc)
//	t.Log("***********************************")
//	t.Log("***********************************")
//	datasDesc := displayDesc(avl)
//	t.Log(datasDesc)
}


func TestAVLTreeSearchRecur(t *testing.T) {
	datas := GetRandNumsJSON("sorted.txt")
	avl := NewAVLTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		avl = avl_insert(avl, datas[i])
	}

	e := AVLTreeSearchRecur(avl, 99267)
	if e {
		t.Log("exist!!!")
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, false)
	}
}

func TestAVLTreeSearchIter(t *testing.T) {
	datas := GetRandNumsJSON("sorted.txt")
	avl := NewAVLTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		avl = avl_insert(avl, datas[i])
	}

	e := AVLTreeSearchIter(avl, 99267)
	if e {
		t.Log("exist!!!")
	} else {
		t.Errorf("Error: expected: %t, actual: %t", true, false)
	}
}

func BenchmarkTreeSearchRecur(b *testing.B) {
	b.StopTimer()
	datas := GetRandNumsJSON("sorted.txt")
	avl := NewAVLTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		avl = avl_insert(avl, datas[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		AVLTreeSearchRecur(avl, 99267)
	}
}

func BenchmarkTreeSearchIter(b *testing.B) {
	b.StopTimer()
	datas := GetRandNumsJSON("sorted.txt")
	avl := NewAVLTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		avl = avl_insert(avl, datas[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		AVLTreeSearchIter(avl, 99267)
	}
}
