package singly_linked_list


type Element struct {
	next *Element
	list * List
	value interface{}
}

func (e *Element) Next() *Element {
	if e.list == nil || e.next == &e.list.root {
		return nil
	}
	return e.next
}

type List struct {
	root Element
	len int
}

func (l *List) init() *List {
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

func (l *List) back() *Element  {
	if l.len == 0 {
		return nil
	}
	p := l.root.next
	for p.next != nil {
		p = p.next
	}

	return p
}

func (l *List) Back() *Element  {
	return l.back()
}

func (l *List) insert(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.next = at.next
	at.next = e

	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{value:v}, at)
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
	return l.insertValue(v, l.back())
}

func (l *List) PushBefore(v interface{}, mark *Element) *Element {
	
}
