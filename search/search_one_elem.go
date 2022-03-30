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

/***************************************AVL树************************************/
type AVLTreeNode struct {
	key		int
	high	int
	left	*AVLTreeNode
	right	*AVLTreeNode
}

func NewAVLTreeNode(value int) *AVLTreeNode {
	return &AVLTreeNode{key: value}
}

func highTree(p *AVLTreeNode) int {
	if p == nil {
		return -1
	} else {
		return p.high
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func left_left_rotation(k *AVLTreeNode) *AVLTreeNode {
	var kl *AVLTreeNode

	kl = k.left
	k.left = kl.right
	kl.right = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kl.high = max(highTree(kl.left), k.high) + 1

	return kl
}

func right_right_rotation(k *AVLTreeNode) *AVLTreeNode {
	var kr *AVLTreeNode

	kr = k.right
	k.right = kr.left
	kr.left = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kr.high = max(k.high, highTree(kr.right)) + 1
	return kr
}


func left_right_rotation(k *AVLTreeNode) *AVLTreeNode {
	k.left = right_right_rotation(k.left)
	
	return left_left_rotation(k)
}

func right_left_rotation(k *AVLTreeNode) *AVLTreeNode {
	k.right = left_left_rotation(k.right)

	return right_right_rotation(k)
}

func avl_insert(avl *AVLTreeNode, key int) *AVLTreeNode {
	if avl == nil {
		avl = NewAVLTreeNode(key)
	} else if key < avl.key {
		avl.left = avl_insert(avl.left, key)
		if highTree(avl.left) - highTree(avl.right) >=2 {
			if key < avl.left.key {
				avl = left_left_rotation(avl)
			} else {
				avl = left_right_rotation(avl)
			}
		}
	} else if key > avl.key {
		avl.right = avl_insert(avl.right, key)
		if highTree(avl.right) - highTree(avl.left) >= 2 {
			if key < avl.right.key {
				avl = right_left_rotation(avl)
			} else {
				avl = right_right_rotation(avl)
			}
		}
	} else if key == avl.key {
		//fmt.Printf("%d has existed!", key)
	}

	avl.high = max(highTree(avl.left), highTree(avl.right)) + 1

	return avl
}


func displayAsc(avl *AVLTreeNode) []int {
	return appendValues([]int{}, avl)
}

func appendValues(values []int, avl *AVLTreeNode) []int {
	if avl != nil {
		values = appendValues(values, avl.left)
		values = append(values, avl.key)
		values = appendValues(values, avl.right)
	}

	return values
}

func displayDesc(avl *AVLTreeNode) []int {
	return appendValuesDesc([]int{}, avl)
}

func appendValuesDesc(values []int, avl *AVLTreeNode) []int {
	if avl != nil {
		values = appendValuesDesc(values, avl.right)
		values = append(values, avl.key)
		values = appendValuesDesc(values, avl.left)
	}
	return values
}

func AVLTreeSearchRecur(avl *AVLTreeNode, key int) bool {
	if avl == nil {
		return false
	} else if key < avl.key {
		return AVLTreeSearchRecur(avl.left, key)
	} else if key > avl.key {
		return AVLTreeSearchRecur(avl.right, key)
	}

	return true
}

func AVLTreeSearchIter(avl *AVLTreeNode, key int) bool {
	var tmp  = avl

	for tmp != nil {
		if key < tmp.key {
			tmp = tmp.left
		} else if key > tmp.key {
			tmp = tmp.right
		} else if key == tmp.key {
			return true
		}
	}

	return false
}


/****************************使用hash表*****************************/

type HashMap map[int]int

func NewMap(nums ...int) HashMap {
	var intMap = make(HashMap)
	for _, v := range nums {
		intMap[v]++
	}

	return intMap
}

func (intmap HashMap) MapSearch(key int) bool {
	if intmap[key] > 0 {
		return true
	}

	return false
}


/*************************自建hash表******************************/

type HashTable []int

func NewHashTable(nums ...int) HashTable {
	var hash [110000]int

	for _, v := range nums {
		hash[v]++
	}
	
	return HashTable(hash[:])
}


func (ht HashTable) HashTableSearch(key int) bool {
	if ht[key] > 0 {
		return true
	}

	return false
}


/*************************红黑树***************************/

func NewRBTree(nums ...int) *rbtree.RBTree {
	rbtree := rbtree.NewRBTree()
	for _, v := range nums {
		rbtree.Insert(int64(v))
	}

	return rbtree
}

