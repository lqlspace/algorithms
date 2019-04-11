/*
*1、需根据elem个数确定Slice的分配空间，而不是先分配空间在init elem，否则如果elem个数少于分配空间，其他空间的零值可能会对BinarySearch算法造成干扰，因为该算法要求数据结构为有序排序
*
*
*/
package search

import (
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

	fmt.Println(si)
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


