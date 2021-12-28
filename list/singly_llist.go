/*
*单向链表：
*（1）插入单个element的时间复杂度为O(1)
*（2）删除单个element的时间复杂度为O(1)
*（3）定位的指定位置的时间复杂度为O(n)
*（4）查找某elment的时间复杂度为
*（5）链表排序的时间复杂度为
*（6）List中存有len，所以返回链表element总数的时间复杂度为O(1)
*（7）定位到头部element的时间复杂度为O(1)
*（8）定位到尾部节点的时间复杂度为O(n)
*（9）移动某个element的时间复杂度为O(1)
*/
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


func (l *List) Remove(val interface{}) {
	if l.len == 0 {
		return
	} else if l.len == 0 {
		if l.front.Value.(int) == val {
			l.front = nil
			l.len--
		}
	}
	prev := l.front
	next := prev.next
	for next != nil {
		if next.Value.(int) == val {
			prev.next = next.next
			l.len--
			break
		}
		prev = next
		next = next.next
	}
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

func (l *List) reverseListRecursive(e *Element) {
	if e == nil {
		return 
	}

	next := e.next
	e.next = l.front
	l.front = e
	l.reverseListRecursive(next)

}

//链表反转方法二：采用递归调用
func (l *List) ReverseListRecursive() {
	if l.len <= 1 {
		return
	}

	front := l.front
	l.reverseListRecursive(front.next)
	front.next = nil
}
