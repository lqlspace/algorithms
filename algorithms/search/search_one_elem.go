/*
*1、需根据elem个数确定Slice的分配空间，而不是先分配空间在init elem，否则如果elem个数少于分配空间，其他空间的零值可能会对BinarySearch算法造成干扰，因为该算法要求数据结构为有序排序
*
*
*/
package search

import (
//	"fmt"
)

type SliceInt []int

//1、采用uint类型，值域为0~2^32-1
func NewSliceInt(nums ...int) SliceInt {
	num := len(nums)
	si := make(SliceInt, num)

	for i := range nums {
		si[i] = nums[i]
	}

	//fmt.Println(si)
	return si
}


func NewSliceIntBitmap(nums ...int) SliceInt {
	var arrInt [400]int

	for _, v := range nums {
		x := v % 32 
		y := v / 32
		arrInt[y] |= (1 << uint(x))
//		fmt.Println(arrInt)
	}

	return SliceInt(arrInt[:])
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


//bitmap结构来查找元素是否存在，时间复杂度O(n)
func (si SliceInt) BitmapSearch(elem int) (exist bool) {
	x := elem % 32
	y := elem / 32
	if si[y] & (1 << uint(x)) > 0 {
		return true
	} else {
		return false
	}
}


