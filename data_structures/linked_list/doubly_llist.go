/*
*双向链表：
*（1）插入单个element时间复杂度O(1)
*（2）删除单个element时间复杂度O(1)
*（3）链表无法直接寻址，故访问特定位置需要从头遍历，时间复杂度为O(n)
*（4）查找特定元素，从头遍历比较查询，时间复杂度O(n)
*（5）排序，待定
*（6）DLList结构中存有len，查询链表element总数的时间复杂度为O(1)
*（7）通过list.root哨兵节点，查找头部节点的时间复杂度为O(1)
*（8）通过list.root哨兵节点，查找尾部节点的时间复杂度为O(1)
*（9）MoveToFront、MoveToBack的时间复杂度为O(1)
*/
package list

type Element struct {
	Value interface{}
	prev *Element
	next *Element
	list *DLList
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type DLList struct {
	root Element
	len int
}


func (l *DLList) Init() *DLList {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func New() *DLList {
	return new(DLList).Init()
}

//通过List结构中存储len，返回链表长度的时间复杂度为O(1)
func (l *DLList) Len() int {
	return l.len
}

//通过哨兵节点l.root，返回头部element的时间复杂度为O(1)
func (l *DLList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}


//通过哨兵节点l.root，返回尾部element的时间复杂度为O(1)
func (l *DLList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

//插入一个element的时间复杂度为O(1)
func (l *DLList) insert(elem, at *Element) *Element {
	elem.prev = at
	elem.next = at.next
	at.next.prev = elem
	at.next = elem
	elem.list = l
	l.len++

	return elem
}

//插入一个value的时间复杂度为O(1)
func (l *DLList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

//在头部插入一个element的时间复杂度为O(1)
func (l *DLList) PushFront(val interface{}) *Element {
	return l.insertValue(val, &l.root)
}

//在尾部插入一个element的时间复杂度为O(1)
func (l *DLList) PushBack(val interface{}) *Element {
	return l.insertValue(val, l.root.prev)
}

//在mark位置插入一个elment的时间复杂度为O(1)，但是定位到mark位置的时间复杂度为O(n)
func (l *DLList) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}


func (l *DLList) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}


//删除1个element的时间复杂度为O(1)
func (l *DLList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--

	return e
}


//删除1个element的时间复杂度为O(1)
func (l *DLList) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

//移动一个element的时间复杂度为O(1)
func (l *DLList) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	at.next.prev = e
	at.next = e

	return e
}


//移动到头部的时间复杂度为O(1)
func (l *DLList) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}

	l.move(e, &l.root)
}


//移动到尾部的时间复杂度为O(1)
func (l *DLList) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

//移动到mark位置之前的时间复杂度为O(1)
func (l *DLList) MoveBefor(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

//移动到mark位置之后的时间复杂度为O(1)
func (l *DLList) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}


func (l *DLList) PushBackList(other *DLList) {
	for e := other.Front(); e != nil; e = e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *DLList) PushFrontList(other *DLList) {
	for e := other.Back(); e != nil; e = e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
