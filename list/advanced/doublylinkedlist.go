package advanced

type Element struct {
	pre *Element
	next *Element
	list *List
	val interface{}
}

func (e *Element) Next() *Element {
	if e.list == nil || e.next == &e.list.root {
		return nil
	}

	return e.next
}

func (e *Element) Pre() *Element {
	if e.list == nil || e.pre == &e.list.root {
		return nil
	}

	return e.pre
}

type List struct {
	root Element
	len int
}

func (l *List) init() *List {
	l.root.pre = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.pre
}

func (l *List) insert(e, at *Element) *Element {
	if e == at {
		return e
	}

	e.pre = at
	e.next = at.next
	e.pre.next = e.next
	e.next.pre = e.pre
	e.list = l
	l.len++

	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insertValue(&Element{val:v}, at)
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.init()
	}
}

func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.pre)
}

func (l *List) PushBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.pre)
}

func (l *List) PushAfter(v interface{}, mark *Element) *Element {
	if mark.list != nil {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.pre.next = e.next
	e.next.pre = e.pre

	e.pre = at
	e.next = at.next
	e.pre.next = e
	e.next.pre = e

	return e
}

func (l *List) MoveToFront(e *Element) {
	if e.list != l || e == l.root.next {
		return
	}
	l.move(e, &l.root)
}

func (l *List) MoveToBack(e *Element) {
	if e.list != l || e == l.root.pre {
		return
	}
	l.move(e, l.root.pre)
}

func (l *List) MoveBefore(e, mark *Element) {
	if e == mark || e.list != l || mark.list != l {
		return
	}

	l.move(e, mark.pre)
}

func (l *List) MoveAfter(e, mark *Element) {
	if e == mark || e.list != l || mark.list != l {
		return
	}

	l.move(e, mark)
}

func (l *List) remove(e *Element) *Element {
	e.pre.next = e.next
	e.next.pre = e.pre
	e.pre = nil
	e.next = nil
	e.list = nil
	l.len--

	return e
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}

	return e.val
}


func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.val, l.root.pre)
	}
}


func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Pre() {
		l.insertValue(e.val, &l.root)
	}
}
