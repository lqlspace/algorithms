/*
*1、需根据elem个数确定Slice的分配空间，而不是先分配空间在init elem，否则如果elem个数少于分配空间，其他空间的零值可能会对BinarySearch算法造成干扰，因为该算法要求数据结构为有序排序
*
*
*/
package search

import (
	"container/list"
	"fmt"
)

type SliceInt []int

//1、采用uint类型，值域为0~2^32-1
func NewSliceInt(nums ...int) SliceInt {
	num := len(nums)
	si := make(SliceInt, num)

	for i := range nums {
		si[i] = nums[i]
	}

	return si
}





func (si SliceInt) IsEmpty() bool {
	return len(si) == 0
}

func (si SliceInt) Num() uint32 {
	return uint32(len(si))
}

func (si SliceInt) binarySearch(elem int, low, high uint32) (idx uint32, exist bool) {
	if high < low {
		return 0, false
	}

	mid := (low + high) / 2

	if elem < si[mid] {
		return si.binarySearch(elem, low, mid-1)
	} else if elem > si[mid] {
		return si.binarySearch(elem, mid+1, high)
	}

	return mid, true
}

//改变原内存空间
func (si SliceInt) BinarySearch(elem int) (index uint32, exist bool) {
	if si.IsEmpty() {
		return 0, false
	}

	low := uint32(0)
	high := si.Num()-1

	return si.binarySearch(elem, low, high)
}

//查找到第一个元素即返回
func (si SliceInt) LineSearch(elem int) (index uint32, exist bool) {
	if si.IsEmpty() {
		return 0, false
	}

	for key, val := range si {
		if elem == val {
			return uint32(key), true
		}
	}

	return 0, false
}



/***********************************链表存储数据********************************************/

//方式一：定义struct包含list.List
type ListInt struct {
	list *list.List
}

func NewListInt(nums ...int) *ListInt {
	l := list.New()
	for _, v := range nums {
		l.PushBack(v)
	}

	return &ListInt {
		list: l,
	}
}

func (li *ListInt) PrintDatas() {
	for e := li.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
}

func (li *ListInt) LineSearch(elem int) bool {
	for e := li.list.Front(); e != nil; e = e.Next() {
		if e.Value == elem {
			return true
		}
	}

	return false
}

//方式二:采用标准库的List，并定义别名然后自定义方法（因为非本地变量无法增加方法），需要注意ListInt并不能直接调用List.List的方法集；
//type ListInt list.List
//
//func NewListInt(nums ...int) *ListInt {
//	l := list.New()
//	for _, v := range nums {
//		l.PushBack(v)
//	}
//
//	return (*ListInt)(l)
//}
//
//func (li *ListInt) PrintDatas() {
//	for e := (*list.List)(li).Front(); e != nil; e = e.Next() {
//		fmt.Printf("%d ", e.Value)
//	}
//}
//
//
//func (li *ListInt) LineSearch(elem int) bool {
//	for e := (*list.List)(li).Front(); e != (*list.List)(li).Back(); e = e.Next() {
//		if e.Value == elem {
//			return true
//		}
//	}
//	if elem == (*list.List)(li).Back().Value {
//		return true
//	}
//
//	return false
//}
//

/*******************************bitmap存储方式************************************/
type Bitmap []int


func NewSliceIntBitmap(nums ...int) Bitmap {
	var arrInt [4000]int

	for _, v := range nums {
		x := v % 32 
		y := v / 32
		arrInt[y] |= (1 << uint(x))
	}

	return Bitmap(arrInt[:])
}


//bitmap结构来查找元素是否存在，时间复杂度O(n)
func (si Bitmap) BitmapSearch(elem int) (exist bool) {
	x := elem % 32
	y := elem / 32
	if si[y] & (1 << uint(x)) > 0 {
		return true
	} else {
		return false
	}
}
