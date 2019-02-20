package list 


type Element struct {
	Value interface{}
	next *Element
	list *List
}


type List struct {
	front *Element
	len int
}

func (l *List) Init() *List {
	l.front = nil
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}


//头部插入的时间复杂度为O(1)
func (l *List) PushFront(val interface{}) *Element {
	e := new(Element)
	e.Value = val
	e.list = l
	e.next = l.front
	l.front = e
	l.len++

	return e
}

//尾部插入的时间复杂度为O(n)
func (l *List) PushBack(val interface{}) *Element {
	e := &Element {
		Value: val,
		list: l,
		next: nil,
	}

	if l.front == nil {
		l.front = e
		l.len++
	} else {
		prev := l.front
		for prev.next != nil {
			prev = prev.next
		}
		prev.next = e
		l.len++
	}
	return e
}

//链表反转方法一：从第2个节点开始到第N个节点，依次进行删除后头部插入；
//用到了两个指针first始终指向ReverseList操作前的第1个节点，e始终执行first的下一个节点；
func (l *List) ReverseList() {
	if l.len <= 1 {
		return 
	}

	first := l.front
	e := first.next
	for e != nil {
		first.next = e.next
		e.next = l.front
		l.front = e
		e = first.next
	}

	first.next = nil
}

//链表反转方法二：采用递归调用

