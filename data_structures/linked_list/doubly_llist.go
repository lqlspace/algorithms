//Doubly Linked List
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

func (l *DLList) Len() int {
	return l.len
}

func (l *DLList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *DLList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *DLList) insert(elem, at *Element) *Element {
	elem.prev = at
	elem.next = at.next
	at.next.prev = elem
	at.next = elem
	elem.list = l
	l.len++

	return elem
}

func (l *DLList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *DLList) PushFront(val interface{}) *Element {
	return l.insertValue(val, &l.root)
}

func (l *DLList) PushBack(val interface{}) *Element {
	return l.insertValue(val, l.root.prev)
}


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


func (l *DLList) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}

	l.move(e, &l.root)
}

func (l *DLList) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

func (l *DLList) MoveBefor(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

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
