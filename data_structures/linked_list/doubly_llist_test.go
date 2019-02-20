package list

import (
	"testing"
	"strconv"
)

func TestPush(t *testing.T) {
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")

	if l.Len() != 3 {
		t.Error("[error] PushFront error, length is wrong !")
	}

	var str string
	for e := l.Front(); e != nil; e = e.Next() {
		str += e.Value.(string)
	}
	t.Log(str)
	if str != "cba" {
		t.Error("[error] Element not right!")
	}

	str = ""
	for e := l.Back(); e != nil; e = e.Prev() {
		str += e.Value.(string)
	}
	if str != "abc" {
		t.Error("[error] Element not right!")
	}
	t.Log("Push element ok: ", str)
}


//拼接两个链表
func TestConcateList(t *testing.T) {
	l1 := New()
	l2 := New()

	l1.PushBack("a")
	l1.PushBack("b")
	l1.PushBack("c")

	l2.PushBack("d")
	l2.PushBack("e")

	l1.PushBackList(l2)
	var str string
	for e := l1.Front(); e != nil; e = e.Next() {
		str += e.Value.(string)
	}

	if str != "abcde" {
		t.Error("[error] concate string error!")
	}
	t.Log("concate string back succeed: ", str)

	l1.PushFrontList(l2)
	str = ""

	for e := l1.Front(); e != nil; e = e.Next() {
		str += e.Value.(string)
	}
	if str != "deabcde" {
		t.Error("[error] concate string error!")
	}
	t.Log("concate string front succeed: ", str)
}


//双向链表实现大数相加
func TestAddNumber(t *testing.T) {
	list1 := New()
	list2 := New()

	list1.PushBack(1)
	list1.PushBack(2)
	list1.PushBack(3)

	list2.PushBack(2)
	list2.PushBack(3)

	var longList *DLList
	var shortList *DLList
	if list1.Len() > list2.Len() {
		longList = list1
		shortList = list2
	} else {
		longList = list2
		shortList = list1
	}
	var carry int
	e1 := longList.Back()
	e2 := shortList.Back()
	for e2 != nil {
		e1.Value = (e1.Value.(int) + e2.Value.(int) + carry) % 10
		carry /= 10
		e1 = e1.Prev()
		e2 = e2.Prev()
	}
	for e1 != nil && carry > 0 {
		e1.Value = (e1.Value.(int) + carry) % 10
		carry /= 10
	}

	if e1 == nil && carry > 0 {
		longList.PushFront(carry)
	}

	var str string
	for e := longList.Front(); e != nil; e = e.Next() {
		str += strconv.Itoa(e.Value.(int))
	}
	if str != "146" {
		t.Error("[error] add operations error")
		return
	}
	t.Log("add operation succeed: ", str)

}
